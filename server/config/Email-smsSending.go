package config

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
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
		fmt.Print(err)
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

	BaseUrl := "https://www.bulksmsnigeria.com/api/v2/sms/create"
	parseUrl, err := url.Parse(BaseUrl)

	if err != nil {
		return false, status.Errorf(codes.Canceled, "unable to resolve request url")
	}

	query := parseUrl.Query()
	query.Set("api_token", BULKAPITOKEN)
	query.Set("from", "Neocloud Technologies")
	query.Set("to", receiverNumber)
	query.Set("body", body)

	parseUrl.RawQuery = query.Encode()

	client := &http.Client{}

	req, err := http.NewRequest("POST", parseUrl.String(), nil)

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
	defer resp.Body.Close()

	return true, nil
}

func SendSmsFunction2(receiverNumber string, body string) (bool, error) {
	EBULKUSERNAMR := os.Getenv("EBULK_USERNAME")
	EBULKAPIKEY := os.Getenv("EBULKAPIKEY")

	BaseUrl := "https://api.ebulksms.com/sendsms"
	parseUrl, err := url.Parse(BaseUrl)

	if err != nil {
		return false, status.Errorf(codes.Canceled, "unable to resolve request url")
	}

	query := parseUrl.Query()
	query.Set("username", EBULKUSERNAMR)
	query.Set("apikey", EBULKAPIKEY)
	query.Set("sender", "NeoCloud")
	query.Set("messagetext", body)
	query.Set("recipients", receiverNumber)

	parseUrl.RawQuery = query.Encode()
	fmt.Println(parseUrl.String())

	resp, err := http.Get(parseUrl.String())

	if err != nil {
		fmt.Println(err)
		return false, status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	fmt.Println(resp)

	return true, nil
}
