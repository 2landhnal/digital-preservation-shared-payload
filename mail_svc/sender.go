package mail_svc

import (
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/logger"
	"github.com/2landhnal/digital-preservation-shared-payload/pb"
	"github.com/2landhnal/digital-preservation-shared-payload/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type GmailSender struct {
	Config *util.Config
	Client pb.NotifierClient
}

func NewGmailSender(config *util.Config) (*GmailSender, error) {
	logger := logger.NewFmtLogger()
	// Connect to the gRPC server
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.NotifyService.ServerAddress, config.NotifyService.GRPCServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not connect to gRPC server: %v", err)
	} else {
		logger.Printf("Connected to gRPC server at %s:%d\n", config.NotifyService.ServerAddress, config.NotifyService.GRPCServerPort)
	}

	client := pb.NewNotifierClient(conn)
	return &GmailSender{
		Config: config,
		Client: client,
	}, nil
}
