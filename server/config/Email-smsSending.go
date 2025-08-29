package config

import (
	"bytes"
	"encoding/json"
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
	BULKAPITOKEN := os.Getenv("BULKAPITOKEN")

	requestUrl := "https://www.bulksmsnigeria.com/api/v2/sms"

	bodyParameters := map[string]string{
		"from": "Neo cloud Technologies",
		"to": receiverNumber,
		"body": body,
		"api_token": BULKAPITOKEN,
		"gateway": "direct-refund",
	}

	jsonBody, err := json.Marshal(bodyParameters)

	if err != nil {
		return false, status.Errorf(codes.Canceled, "an error occured while processing your request %v", err)
	}

	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonBody))

	// client := &http.Client{}

	// req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonBody))

	// if err != nil {
	// 	return false, status.Errorf(codes.Canceled, "an error occured while processing your request %v", err)
	// }

	// req.Header.Set("Accept", "application/json")
	// req.Header.Set("Content-Type", "application/json")

	// resp, err := client.Do(req)

	if err != nil {
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	defer resp.Body.Close()

	return true, nil
}


func SendSmsFunction2(receiverNumber string, body string)(bool, error){
	EBULK_USERNAMR := os.Getenv("EBULK_USERNAME")
	EBULKAPIKEY := os.Getenv("EBULKAPIKEY")

	requestUrl := "https://api.ebulksms.com/sendsms.json"

	bodyParameters := map[string]interface{}{
		"sms": map[string]interface{}{
			"auth": map[string]interface{}{
				"username": EBULK_USERNAMR,
				"apikey": EBULKAPIKEY,
			},
			"message": map[string]interface{}{
				"sender": "Neo Cloud Technologies",
				"messagetext": body,
				"flash": "0",
			},
			"recipients": map[string]interface{}{
				"gsm":map[string]interface{}{
					"msidn": receiverNumber,
					"msgid": "",
				},
			},
			"dndsender": 1,
		},
	}

	jsonBody, err := json.MarshalIndent(bodyParameters, "", "")

	if err != nil {
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}
	defer resp.Body.Close()

	return true, nil
}
