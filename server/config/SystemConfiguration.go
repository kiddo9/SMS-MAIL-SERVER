package config

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
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
	
	//create message
	m := gomail.NewMessage()
	m.SetHeader("From", MailVariables["SMTP_USER"])
	m.SetHeader("To", email)
	m.SetHeader("Subject", "OTP Verification")
	m.SetBody("text/html", body.String())

	smsPortConvertion, err := strconv.Atoi(MailVariables["SMTP_PORT"])

	if err != nil {
		fmt.Print(err, MailVariables)
		return false, status.Errorf(codes.Internal, "internal server error")
	}

	d := gomail.NewDialer(MailVariables["SMTP_HOST"], int(smsPortConvertion), MailVariables["SMTP_USER"], MailVariables["SMTP_PASSWORD"])

	err = d.DialAndSend(m)

	if err != nil{
	return false, status.Errorf(codes.Internal, "internal server error")
	}

	return true, nil
}