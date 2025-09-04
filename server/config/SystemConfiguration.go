package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/joho/godotenv"
	"github.com/kiddo9/SMS-MAIL-SERVER/structures"
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

func LoadTemplate(templateType string, id int)(string, error){
	var TemplateContent string 

	_, err := os.Open("storage/Templates.json")

	if err != nil {
		return "error", status.Errorf(codes.Internal, "Internal Server Error")
	}

	data, err := os.ReadFile("storage/Templates.json")

	if err != nil {
		return "error", status.Errorf(codes.Internal, "Internal Server Error")
	}

	var Array structures.TemplateStruct
	err = json.Unmarshal(data, &Array)

	if err != nil {
		return "error", status.Errorf(codes.Internal, "Internal Server Error")
	}

	if templateType == "email" {
		EmailTemplateList := Array.Templates.EmailsTemp

		for _, emailTemp := range EmailTemplateList {
			if emailTemp.ID == id {
				TemplateContent = emailTemp.TemplateContent
				return TemplateContent, nil
			}
		}
	}else if templateType == "sms" {
		SmsTemplateList := Array.Templates.SmsTemp

		for _, smsTemp := range SmsTemplateList {
			if smsTemp.ID == id {
				 TemplateContent = smsTemp.TemplateContent
				return TemplateContent, nil
			}
		}
	}else{
		return "error", status.Errorf(codes.InvalidArgument, "Invalid template type")
	}

	return "", nil
}

func AuthenticationMailling(email string, otp string) (bool, error) {
	tmpl, err := template.New("otpEmail").Parse(templates.OTPEmail)
	if err != nil {
		panic(err)
	}

	info := map[string]string{
		"Name":          "Admin or Oracle",
		"EmailContent":  "This is your one time password and will expire in 3 minutes",
		"OTP":           otp,
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

func BulkEmail(name string, pendingPrice string, course string, Date string, emailAddress string, phoneNumber string, senderEmail string, tempType string, tempId int) (bool, error) {
	 content, err := LoadTemplate(tempType, tempId)
	 fmt.Println(content)

	 if err != nil || content == "error" {
		fmt.Println(err)
		return false, status.Errorf(codes.Internal, "could not resolve template %v", err)
	 }

	tmp, err := template.New("BatchUploadEmail").Parse(content)

	if err != nil {
		return false, status.Errorf(codes.Internal, "could not parse template %v", err)
	}

	price := fmt.Sprintf("%v", pendingPrice)
	number := fmt.Sprintf("%v", phoneNumber)

	Datas := map[string]string{
		"Name":         name,
		"PendingPrice": price,
		"course":       course,
		"Date":         Date,
		"phoneNumber":  number,
		"EmailAddress": senderEmail,
	}

	var body bytes.Buffer
	err = tmp.Execute(&body, Datas)
	if err != nil {
		return false, status.Errorf(codes.Aborted, "process could not be completed. %v", err)
	}

	_, err = sendEmail(emailAddress, body, "Friendly Reminder")
	if err != nil {
		return false, status.Errorf(codes.Aborted, "error occured will processing the bulk email. %v", err)
	}

	return true, nil
}

func BulkSms(name string, pendingPrice string, course string, Date string, phoneNumber string, senderEmail string, receverNumber string, method string) (bool, error) {
	tmp, err := template.New("BatchUploadSMS").Parse(templates.SmsTemp)

	if err != nil {
		panic(err)
	}

	price := fmt.Sprintf("%v", pendingPrice)
	number := fmt.Sprintf("%v", phoneNumber)

	Datas := map[string]string{
		"Name":         name,
		"PendingPrice": price,
		"course":       course,
		"Date":         Date,
		"phoneNumber":  number,
		"EmailAddress": senderEmail,
	}

	var body bytes.Buffer
	err = tmp.Execute(&body, Datas)
	if err != nil {
		return false, status.Errorf(codes.Aborted, "process could not be completed")
	}

	if method == "Bulksms" {
		resp, err := SendSmsUsingBulk(receverNumber, body.String())
		if err != nil {
			fmt.Println("seems the first sms server is down. moving to the second sever")
			return false, status.Errorf(codes.Aborted, "process could not be completed %v", err)
		}
		fmt.Println(resp)
	}

	if method == "EBulksms" {
		_, err := SendSmsFunction2(receverNumber, body.String())

		if err != nil {
			return false, status.Errorf(codes.Aborted, "process could not be completed. seems both sms servers are down")
		}
	}

	return true, nil
}
