package config

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/twilio/twilio-go"
	openApi "github.com/twilio/twilio-go/rest/api/v2010"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func sendEmail(email string, body bytes.Buffer, subject string)(bool, error){
	m := gomail.NewMessage()
	m.SetHeader("From", MailVariables["SMTP_USER"])
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	smsPortConvertion, err := strconv.Atoi(MailVariables["SMTP_PORT"])

	if err != nil {
		fmt.Print(err, MailVariables)
		return false, status.Errorf(codes.Internal, "internal server error")
	}

	d := gomail.NewDialer(MailVariables["SMTP_HOST"], int(smsPortConvertion), MailVariables["SMTP_USER"], MailVariables["SMTP_PASSWORD"])

	err = d.DialAndSend(m)

	if err != nil{
		fmt.Println(err)
	return false, status.Errorf(codes.Internal, "internal server error %v", err)
	}

	return true, nil
}

func SendSmsUsingTwiilo( receiverNumber string, body string)(bool, error){
	senderNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	twilioAccountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	twilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")

	smsClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioAccountSID,
		Password: twilioAuthToken,
	})

	//create message 
	createmessageParams := &openApi.CreateMessageParams{}
	createmessageParams.SetTo(receiverNumber)
	createmessageParams.SetFrom(senderNumber)
	createmessageParams.SetBody(body)

	//send sms
	_, err := smsClient.Api.CreateMessage(createmessageParams)

	if err != nil {
		fmt.Println(err)
		return false, status.Errorf(codes.Internal, "Something went wrong %v", err)
	}

	return true, nil
}