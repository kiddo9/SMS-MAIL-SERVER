package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	pb "github.com/kiddo9/SMS-MAIL-SERVER/message/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Wallet struct {
	pb.UnimplementedSmsServicesServer
}

func (w *Wallet) EbulkSmsWallet(req *pb.EbulkSms, stream pb.SmsServices_EbulkSmsWalletServer) error {
	EBULKUSERNAME := os.Getenv("EBULK_USERNAME")
	EBULKAPIKEY := os.Getenv("EBULKAPIKEY")

	requestUrl := fmt.Sprintf("https://api.ebulksms.com/balance/%v/%v", EBULKUSERNAME, EBULKAPIKEY)

	resp, err := http.Get(requestUrl)
	
	if err != nil {
		return status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return status.Errorf(codes.Canceled, "unable to read body %v", err)
	}
	
	return stream.Send(&pb.EbulkSmsResponse{
		Response: string(body),
	})
}

func (w *Wallet) BulkSmsWallet(req *pb.BulkSms, stream pb.SmsServices_BulkSmsWalletServer) error {
	BULKAPITOKEN := os.Getenv("BULKAPITOKEN")

	requestUrl := fmt.Sprintf("https://www.bulksmsnigeria.com/api/v2/balance?api_token=%v", BULKAPITOKEN)
	fmt.Println(requestUrl)

	resp, err := http.Get(requestUrl)
	
	if err != nil {
		return status.Errorf(codes.Canceled, "unable to complete your request %v", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return status.Errorf(codes.Canceled, "unable to read body %v", err)
	}
	
	return stream.Send(&pb.BulkSmsResponse{
		Response: string(body),
	})
	
}