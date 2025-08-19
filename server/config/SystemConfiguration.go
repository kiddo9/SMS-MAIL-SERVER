package config

import (
	"os"
	"strconv"

	"github.com/go-gomail/gomail"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


var SMTPsHOST string = os.Getenv("EMAIL_HOST")
var SMTPPORT  int 
var SMTPUSER  string = os.Getenv("EMAIL_USERNAME")
var SMTPPASSWORD  string= os.Getenv("EMAIL_PASSWORD")

var SmsVariables = map[string]string{
	"SMS_API_KEY": os.Getenv(""),
	"SMS_API_URL": os.Getenv(""),
}

func AuthenticationMailling(email string, otp string) (bool, error){
	

	//create message
	m := gomail.NewMessage()
	m.SetHeader("From", SMTPUSER)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "OTP Verification")
	m.SetBody("text/html", "<p>Dear User,</p><p>Your OTP for verification is: <strong>" + otp + "</strong></p><p>This otp last for 3 minutes </p>")

	smsPortConvertion, err := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	if err != nil {
		return false, status.Errorf(codes.Internal, "internal server error")
	}

	d := gomail.NewDialer(SMTPsHOST, int(smsPortConvertion), SMTPUSER, SMTPPASSWORD)

	err = d.DialAndSend(m)

	if err != nil{
	return false, status.Errorf(codes.Internal, "internal server error")
	}

	return true, nil
}