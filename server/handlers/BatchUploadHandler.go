package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"github.com/xuri/excelize/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FileUploadStruct struct {
	pb.UnimplementedFileUploadServicesServer
}

func (f *FileUploadStruct) FileUpload(ctx context.Context, req *pb.FileUploadRequest)(*pb.FileUploadResponse, error){

	file := req.GetContent()

	readFile, err := excelize.OpenReader(io.NopCloser(bytes.NewReader(file)))

	if err != nil {
		return nil, status.Errorf(codes.Unknown, "file can't be opened %v", err)
	}

	fmt.Println(readFile)

	rows, err := readFile.GetRows("sheet1")

	if err != nil {
		return nil, status.Errorf(codes.Unknown, "file can't be read. make sure file sent is excel %v", err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
	return &pb.FileUploadResponse{
		Status: true,
		Message: "email and sms sent",
	}, nil
}