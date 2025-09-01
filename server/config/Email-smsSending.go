package config

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func sendEmail(email string, body bytes.Buffer, subject string) (bool, error) {
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

	if err != nil {
		fmt.Println(err)
		return false, status.Errorf(codes.Internal, "internal server error %v", err)
	}

	return true, nil
}

func SendSmsUsingBulk(receiverNumber string, body string) (bool, error) {
	BULKAPITOKEN := os.Getenv("BULKAPITOKEN")

	requestUrl := fmt.Sprintf("https://www.bulksmsnigeria.com/api/v2/sms/create?api_token=%v&from=%v&to=%v&body=%v", BULKAPITOKEN, "Neocloud_Technologies", receiverNumber, body)

	client := &http.Client{}

	req, err := http.NewRequest("POST", requestUrl, nil)

	if err != nil {
		fmt.Println(err)
		return false, status.Errorf(codes.Canceled, "an error occured while processing your request %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	fmt.Println(resp, resp.Status)
	defer resp.Body.Close()

	return true, nil
}

func SendSmsFunction2(receiverNumber string, body string) (bool, error) {
	EBULKUSERNAMR := os.Getenv("EBULK_USERNAME")
	EBULKAPIKEY := os.Getenv("EBULKAPIKEY")

	requestUrl := fmt.Sprintf("https://api.ebulksms.com/sendsms?username=%v&apikey=%v&sender=%v&messagetext=%v&flash=0&recipients=%v", EBULKUSERNAMR, EBULKAPIKEY, "NeoCloud", body, receiverNumber)

	_, err := http.Get(requestUrl)
	
	if err != nil {
		fmt.Println(err)
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	return true, nil
}
