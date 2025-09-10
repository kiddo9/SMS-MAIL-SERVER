package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/kiddo9/SMS-MAIL-SERVER/config"
	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"github.com/kiddo9/SMS-MAIL-SERVER/structures"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type FileUploadStruct struct {
	pb.UnimplementedFileUploadServicesServer
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func (f *FileUploadStruct) FileUpload(ctx context.Context, req *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	var datas []byte
	var body []structures.AdminStructs
	var Admin structures.AdminStructs

	_, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	datas, err = os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(datas, &body)
	if err != nil {
		panic(err)
	}

	Admin = body[0]

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	var messageMethod string
	var MMth string
	var EmailId []string = md["emailid"]
	var smsId []string = md["smsid"]
	var emailIdStr string
	var smsIdStr string
	var Message string
	var Id int
	var SmsId int

	// if len(md["send_using"]) == 0 {
	// 	return nil, status.Errorf(codes.InvalidArgument, "missing argument")
	// }

	if len(EmailId) == 0 && len(smsId) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "missing emailId or smsId")
	}

	if contains(md["send_using"], "email") && (contains(md["send_using"], "Bulksms") || contains(md["send_using"], "EBulksms")) {
		emailIdStr = EmailId[0]
		smsIdStr = smsId[0]

		Id, err = strconv.Atoi(emailIdStr)

		if err != nil {
			return nil, err
		}

		SmsId, err = strconv.Atoi(smsIdStr)

		if err != nil {
			return nil, err
		}
	} else if contains(md["send_using"], "email") {
		emailIdStr = EmailId[0]

		Id, err = strconv.Atoi(emailIdStr)

		if err != nil {
			return nil, err
		}
	} else if contains(md["send_using"], "Bulksms") || contains(md["send_using"], "EBulksms") {
		smsIdStr = smsId[0]

		SmsId, err = strconv.Atoi(smsIdStr)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
	}

	file := req.GetContent()
	data := req.GetDate()

	readFile, err := excelize.OpenReader(io.NopCloser(bytes.NewReader(file)))

	if err != nil {
		return nil, status.Errorf(codes.Unknown, "file can't be opened %v", err)
	}

	result := make(map[string][][]string)

	sheets := readFile.GetSheetList()

	for _, sheet := range sheets {
		rows, err := readFile.GetRows(sheet)

		if err != nil {
			return nil, status.Errorf(codes.Unknown, "file can't be read. make sure file sent is excel %v", err)
		}

		result[sheet] = rows

		for idx, row := range result[sheet] {

			if idx == 0 {
				continue
			}

			name := row[0]
			pendingPrice := row[12]
			phone := row[3]
			course := row[2]
			email := row[5]

			if strings.TrimSpace(pendingPrice) != "" {
				if contains(md["send_using"], "email") && (contains(md["send_using"], "Bulksms") || contains(md["send_using"], "EBulksms")) {
					if contains(md["send_using"], "EBulksms") {
						MMth = "EBulksms"
					} else {
						MMth = "Bulksms"
					}
					_, err := config.BulkEmail(name, pendingPrice, course, data, email, Admin.Phone, Admin.Email, "email", Id)

					if err != nil {
						return nil, status.Errorf(codes.Unknown, "unable to complete bulk email")
					}

					_, err = config.BulkSms(name, pendingPrice, course, data, Admin.Phone, Admin.Email, phone, MMth, "sms", SmsId)

					if err != nil {
						return nil, status.Errorf(codes.Unknown, "unable to complete sms email")
					}

					Message = "email and sms sent"
				} else {
					for _, method := range md["send_using"] {
						if method != "email" && method != "Bulksms" && method != "EBulksms" {
							return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
						}

						messageMethod = method
					}

					if messageMethod == "email" {
						_, err = config.BulkEmail(name, pendingPrice, course, data, email, Admin.Phone, Admin.Email, messageMethod, Id)

						if err != nil {
							return nil, status.Errorf(codes.Unknown, "unable to complete bulk email")
						}

						Message = "email sent"
					}

					if messageMethod == "Bulksms" {
						_, err = config.BulkSms(name, pendingPrice, course, data, Admin.Phone, Admin.Email, phone, messageMethod, "sms", SmsId)

						if err != nil {
							return nil, status.Errorf(codes.Unknown, "unable to complete sms email")
						}

						Message = "sms sent"
					}

					if messageMethod == "EBulksms" {
						_, err = config.BulkSms(name, pendingPrice, course, data, Admin.Phone, Admin.Email, phone, messageMethod, "sms", SmsId)

						if err != nil {
							return nil, status.Errorf(codes.Unknown, "unable to complete sms email")
						}

						Message = "sms sent"
					}
				}
			}
		}
	}

	return &pb.FileUploadResponse{
		Status:  true,
		Message: Message,
	}, nil
}
