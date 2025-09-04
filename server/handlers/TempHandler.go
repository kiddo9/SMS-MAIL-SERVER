package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"github.com/kiddo9/SMS-MAIL-SERVER/structures"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Temp struct {
	pb.UnimplementedTemplateServicesServer
}

var templateFileName string = "storage/Templates.json"
var templateData structures.TemplateStruct

func LoadTemplate() error {
	_, err := os.Open(templateFileName)

	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(templateFileName)

	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}

	err = json.Unmarshal(file, &templateData)

	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}

	return nil
}

func (t *Temp) CreateEmailTemplate(ctx context.Context, req *pb.TemplateRequest)(*pb.Response, error) {
	err := LoadTemplate()

	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	templateReqBody := structures.EmailTemp{
		ID:              len(templateData.Templates.EmailsTemp) + 1,
		TemplateName:    req.GetTemplateName(),
		TemplateContent: req.GetTemplateContent(),
		Date:            req.GetDate(),
	}

	emailsTemp := templateData.Templates.EmailsTemp
	emailsTemp = append(emailsTemp, templateReqBody)
	templateData.Templates.EmailsTemp = emailsTemp

	updatedData, err := json.MarshalIndent(templateData, " ", " ")

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	err = os.WriteFile(templateFileName, updatedData, 0644)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &pb.Response{Status: true, Message: "Email Template Created Successfully"}, nil
}

func (t *Temp) CreateSmsTemplate(ctx context.Context, req *pb.SmsTemplateRequest)(*pb.Response, error){
	err := LoadTemplate()

	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	smsTemplateReqBody := structures.SmsTemp{
		ID: len(templateData.Templates.SmsTemp) + 1,
		TemplateName: req.SmsTemplateName,
		TemplateContent: req.SmsTemplateContent,
		Date: req.Date,
	}

	smsTemp := templateData.Templates.SmsTemp
	smsTemp =  append(smsTemp, smsTemplateReqBody)
	templateData.Templates.SmsTemp = smsTemp

	updatedData, err := json.MarshalIndent(templateData, " ", " ")

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	err = os.WriteFile(templateFileName, updatedData, 0644)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	return &pb.Response{Status: true, Message: "sms Template Created Successfully"}, nil
}