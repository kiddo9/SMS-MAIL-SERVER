package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/kiddo9/SMS-MAIL-SERVER/config"
	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"github.com/kiddo9/SMS-MAIL-SERVER/structures"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileUploadStruct struct {
	pb.UnimplementedFileUploadServicesServer
}

var datas []byte
var body []structures.AdminStructs
var Admin structures.AdminStructs

func init() {
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

	body[0] = Admin
}

func (f *FileUploadStruct) FileUpload(ctx context.Context, req *pb.FileUploadRequest)(*pb.FileUploadResponse, error){

	file := req.GetContent()

	readFile, err := excelize.OpenReader(io.NopCloser(bytes.NewReader(file)))

	if err != nil {
		return nil, status.Errorf(codes.Unknown, "file can't be opened %v", err)
	}

	result := make(map[string][][]string)

	sheets := readFile.GetSheetList()

	for _, sheet := range sheets{
		rows, err := readFile.GetRows(sheet)

		if err != nil {
			return nil, status.Errorf(codes.Unknown, "file can't be read. make sure file sent is excel %v", err)
		}

		result[sheet] = rows

		for _, row := range result[sheet] {
			name := row[0]
			pendingPrice := row[1]
			course := row[2]
			data := row[3]
			email := row[4]
			
			_, err = config.BulkEmail(name, pendingPrice, course, data, Admin.Email, Admin.Phone, email)

			if err != nil {
				return nil, status.Errorf(codes.Unknown, "unable to complete bulk email and sms")
			}
		}
	}

	return &pb.FileUploadResponse{
		Status: true,
		Message: "email and sms sent",
	}, nil
}