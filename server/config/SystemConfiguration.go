package config

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"github.com/kiddo9/SMS-MAIL-SERVER/templates"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


var MailVariables map[string]string
var SmsVariables map[string]string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	MailVariables = map[string]string{
		"SMTP_HOST":     os.Getenv("EMAIL_HOST"),
		"SMTP_PORT":     os.Getenv("EMAIL_PORT"),
		"SMTP_USER":     os.Getenv("EMAIL_USERNAME"),
		"SMTP_PASSWORD": os.Getenv("EMAIL_PASSWORD"),
	}

		SmsVariables = map[string]string{
		"SMS_API_KEY": os.Getenv(""),
		"SMS_API_URL": os.Getenv(""),
	}
}


func AuthenticationMailling(email string, otp string) (bool, error){
	tmpl, err := template.New("otpEmail").Parse(templates.OTPEmail)
    if err != nil {
        panic(err)
    }

	info := map[string]string{
		"Name": "Admin or Oracle",
		"EmailContent": "This is your one time password and will expire in 3 minutes",
		"OTP": otp,
		"ExpiryMinutes": "3",
	}

	 var body bytes.Buffer
    if err := tmpl.Execute(&body, info); err != nil {
        panic(err)
    }

	_, err = sendEmail(email, body, "OTP Verification")
	if err != nil {
		panic(err)
	}

	return true, nil
}


func BulkEmail(name string, pendingPrice string, course string, Date string, emailAddress string, phoneNumber string, senderEmail string)(bool, error){
	tmp, err := template.New("BatchUploadEmail").Parse(templates.EmailTemplate1)

	if err != nil {
		panic(err)
	}

	price := fmt.Sprintf("%v", pendingPrice)
	number := fmt.Sprintf("%v", phoneNumber)

	Datas := map[string]string{
		"Name": name,
		"PendingPrice": price,
		"course": course,
		"Date": Date,
		"phoneNumber": number,
		"EmailAddress": emailAddress,
	}

	var body bytes.Buffer
	err = tmp.Execute(&body, Datas)
	if err != nil {
		return false, status.Errorf(codes.Aborted, "process could not be completed")
	}

	_, err = sendEmail(senderEmail, body, "Friendly Reminder")
	if err != nil {
		panic(err)
	}

	return true, nil
}