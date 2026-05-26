package handlers

import (
	"bytes"
	"context"
	"io"
	"log"
	"strings"
	"sync"

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

// goroutine function
func messageWorker(
	jobs <-chan structures.MessageJob,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for job := range jobs {

		sendEmail := contains(job.Method, "email")
		sendSms := contains(job.Method, "Bulksms") || contains(job.Method, "EBulksms")

		MMth := ""
		if contains(job.Method, "EBulksms") {
			MMth = "EBulksms"
		} else if contains(job.Method, "Bulksms") {
			MMth = "Bulksms"
		}

		if sendEmail {
			_, err := config.BulkEmail(
				job.Name,
				job.PendingPrice,
				job.Course,
				job.Data,
				job.Email,
				job.Admin.Phone,
				job.Admin.Email,
				"email",
				job.EmailId,
			)

			if err != nil {
				log.Println("email failed:", err)
				continue
			}
		}

		if sendSms {
			_, err := config.BulkSms(
				job.Name,
				job.PendingPrice,
				job.Course,
				job.Data,
				job.Admin.Phone,
				job.Admin.Email,
				job.Phone,
				MMth,
				"sms",
				job.SmsId,
			)

			if err != nil {
				log.Println("sms failed:", err)
				continue
			}
		}
	}
}

func (f *FileUploadStruct) FileUpload(ctx context.Context, req *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	var Message string

	jobs := make(chan structures.MessageJob, 100)
	wg := sync.WaitGroup{}

	numWorkers := 10

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go messageWorker(jobs, &wg)
	}

	file := req.GetContent()
	data := req.GetDate()
	EmailId := req.GetEmailId()
	SmsId := req.GetSmsId()

	var requiredColumns = map[string]bool{
		"name":         false,
		"course":       false,
		"phone":        false,
		"email":        false,
		"pendingprice": false,
	}

	readFile, err := excelize.OpenReader(io.NopCloser(bytes.NewReader(file)))

	if err != nil {
		return nil, status.Errorf(codes.Unknown, "file can't be opened %v", err)
	}

	//result := make(map[string][][]string)
	headerIndex := make(map[string]int)

	sheets := readFile.GetSheetList()

	for _, sheet := range sheets {
		rows, err := readFile.GetRows(sheet)

		if err != nil {
			return nil, status.Errorf(codes.Unknown, "file can't be read. make sure file sent is excel %v", err)
		}

		//result[sheet] = rows
		headerRow := rows[0]

		for idx, header := range headerRow {
			h := strings.ToLower(strings.TrimSpace(header))
			if _, exists := requiredColumns[h]; exists {
				headerIndex[h] = idx
				requiredColumns[h] = true
			}
		}

		for col, found := range requiredColumns {
			if !found {
				return nil, status.Errorf(
					codes.InvalidArgument,
					"missing required column: %s",
					col,
				)
			}
		}

		for idx, row := range rows {

			if idx == 0 {
				continue
			}

			if strings.TrimSpace(row[headerIndex["pendingprice"]]) == "" {
				continue
			}

			job := structures.MessageJob{
				Name:         row[headerIndex["name"]],
				PendingPrice: row[headerIndex["pendingprice"]],
				Course:       row[headerIndex["course"]],
				Phone:        row[headerIndex["phone"]],
				Email:        row[headerIndex["email"]],
				Data:         data,
				Method:       md["x-send-using"],
				EmailId:      int(EmailId),
				SmsId:        int(SmsId),
				//Admin:        Admin,
			}

			jobs <- job
		}
	}

	close(jobs)
	wg.Wait()

	return &pb.FileUploadResponse{
		Status:  true,
		Message: Message,
	}, nil
}
