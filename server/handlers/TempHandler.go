package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

func (t *Temp) GetEmailTemplateById(ctx context.Context, req *pb.GetATemplateRequest)(*pb.GetATemplateResponse, error){
	err := LoadTemplate()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	EmailTempId := req.GetId()
	convertToNumber, err := strconv.Atoi(EmailTempId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	emailsTemp := templateData.Templates.EmailsTemp
	
	for _, v := range emailsTemp {
		if v.ID == convertToNumber {
			convertToString := strconv.Itoa(v.ID)

			return &pb.GetATemplateResponse{
				Template: &pb.Template{
					Id:              convertToString,
					TemplateName:    v.TemplateName,
					TemplateContent: v.TemplateContent,
					Date:            v.Date,
				},
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Email Template Not Found")
}


func (t *Temp) GetSmsTemplateById(ctx context.Context, req *pb.GetAnSmsTemplateRequest)(*pb.GetAnSmsTemplateResponse, error){
	err := LoadTemplate()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	SmsTempId := req.GetId()
	convertToNumber, err := strconv.Atoi(SmsTempId)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	smsTemp := templateData.Templates.SmsTemp
	
	for _, v := range smsTemp {
		if v.ID == convertToNumber {
			convertToString := strconv.Itoa(v.ID)

			return &pb.GetAnSmsTemplateResponse{
				SmsTemplate: &pb.SmsTemplate{
					Id: convertToString,
					SmsTemplateName:    v.TemplateName,
					SmsTemplateContent: v.TemplateContent,
					Date: v.Date,
				},
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Email Template Not Found")
}

func (t *Temp) EditEmailTemplate(ctx context.Context, req *pb.TemplateEditRequest)(*pb.Response, error){
	err := LoadTemplate()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	EmailTempBody := req.Template
	convertToNumber, err :=  strconv.Atoi(EmailTempBody.Id)

	if err != nil{
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	updatedEmailTemp := structures.EmailTemp{
		ID: convertToNumber,
		TemplateName: EmailTempBody.TemplateName,
		TemplateContent: EmailTempBody.TemplateContent,
		Date: EmailTempBody.Date,
	}

	emailsTemp := templateData.Templates.EmailsTemp
	
	for idx, v := range emailsTemp {
		if v.ID == convertToNumber {
			v.TemplateName = updatedEmailTemp.TemplateName
			v.TemplateContent = updatedEmailTemp.TemplateContent
			v.Date = updatedEmailTemp.Date

			emailsTemp[idx] = v		
			
			updatedData, err := json.MarshalIndent(templateData, " ", " ")

			if err != nil {
				return nil, err
			}

			err = os.WriteFile(templateFileName, updatedData, 0644)

			if err != nil {
				return nil, err
			}

			return &pb.Response{
				Status: true,
				Message: "Email template Edited successfull",
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Email Template Not Found")
}

func (t *Temp) EditSmsTemplate(ctx context.Context, req *pb.SmsTemplateEditRequest)(*pb.Response, error){
	err := LoadTemplate()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}

	smsTempBody := req.Smstemplate
	convertToNumber, err :=  strconv.Atoi(smsTempBody.Id)

	if err != nil{
		return nil, status.Errorf(codes.Internal, "Internal Server Error")
	}
	updatedSmsTemp := structures.SmsTemp{
		ID: convertToNumber,
		TemplateName: smsTempBody.SmsTemplateName,
		TemplateContent: smsTempBody.SmsTemplateContent,
		Date: smsTempBody.Date,
	}

	smsTemp := templateData.Templates.SmsTemp
	
	for idx, v := range smsTemp {
		if v.ID == convertToNumber {
			v.ID = updatedSmsTemp.ID
			v.TemplateName = updatedSmsTemp.TemplateName
			v.TemplateContent = updatedSmsTemp.TemplateContent
			v.Date = updatedSmsTemp.Date

			smsTemp[idx] = v		
			
			updatedData, err := json.MarshalIndent(templateData, " ", " ")

			if err != nil {
				return nil, err
			}

			err = os.WriteFile(templateFileName, updatedData, 0644)

			if err != nil {
				return nil, err
			}

			return &pb.Response{
				Status: true,
				Message: "Email template Edited successfull",
			}, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "Email Template Not Found")
}

func (t *Temp) AllTemplates(req *pb.TemplateFetchRequest, stream pb.TemplateServices_AllTemplatesServer) error {
	err := LoadTemplate()

	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}

	// Stream all email templates
	for _, email := range templateData.Templates.EmailsTemp {
		resp := &pb.TemplateFetchResponse{
				EmailTemplate: &pb.Template{
				Id:              strconv.Itoa(email.ID),
				TemplateName:    email.TemplateName,
				TemplateContent: email.TemplateContent,
				Date:            email.Date,
			},
		}
		
		if err := stream.Send(resp); err != nil {
			return status.Errorf(codes.Internal, "Failed to stream email template: %v", err)
		}
	}

	// Stream all SMS templates
	for _, sms := range templateData.Templates.SmsTemp {
		resp := &pb.TemplateFetchResponse{
				SmsTemplate: &pb.SmsTemplate{
				Id:                strconv.Itoa(sms.ID),
				SmsTemplateName:   sms.TemplateName,
				SmsTemplateContent:sms.TemplateContent,
				Date:              sms.Date,
			},
		}
		if err := stream.Send(resp); err != nil {
			return status.Errorf(codes.Internal, "Failed to stream SMS template: %v", err)
		}
	}

	return nil
}