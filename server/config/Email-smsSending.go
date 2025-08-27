package config

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
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

func SendSmsUsingBulk( receiverNumber string, body string)(bool, error){
	BULK_USERNAMR := os.Getenv("BULK_USERNAMR")
	BULKPASSWORD := os.Getenv("BULKPASSWORD")

	requestUrl :=fmt.Sprintf("https://api.bulksms.com/v1/messages/send?to=%v&body=%v", receiverNumber, body) 

	client := &http.Client{}

	req, err := http.NewRequest("GET", requestUrl, nil)

	if err != nil {
		return false, status.Errorf(codes.Canceled, "an error occured while processing your request %v", err)
	}

	credentials := fmt.Sprintf("%v:%v", BULK_USERNAMR, BULKPASSWORD)
	encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+ encodedCredentials)

	resp, err := client.Do(req)

	if err != nil {
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	defer resp.Body.Close()

	return true, nil
}


func SendSmsFunction2( receiverNumber string, body string)(bool, error){
	EBULK_USERNAMR := os.Getenv("EBULK_USERNAMR")
	EBULKAPIKEY := os.Getenv("EBULKAPIKEY")

	requestUrl := fmt.Sprintf("https://api.ebulksms.com/sendsms?username=%v&apikey=%v&sender=%v&message=%v&flash=0&recipients=%v&dndsender=%v",EBULK_USERNAMR,EBULKAPIKEY,"Neo Cloud Technologies",body,receiverNumber, 0 )

	resp, err := http.Get(requestUrl)
	if err != nil {
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}
	defer resp.Body.Close()

	return true, nil
}

//?username=your_email_address& apikey=your_apikey&sender=your_sender_name& messagetext=your_message&flash=0& recipients=23480...,23470...&dndsender=1