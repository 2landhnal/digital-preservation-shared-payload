package mail_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/logger"
	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (gmailSender *GmailSender) SendEmail(ctx context.Context, req *pb.SendMailRequest) (*pb.SendMailResponse, error) {
	// Call the Identify RPC
	logger := logger.NewFmtLogger()
	logger.Println("Sending email to:", req.To)
	resp, err := gmailSender.Client.SendEmail(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call SendEmail GAPI: %v", err)
	}

	return resp, nil
}
