package collect_svc

import (
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/logger"
	"github.com/2landhnal/digital-preservation-shared-payload/pb"
	"github.com/2landhnal/digital-preservation-shared-payload/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CollectServicer struct {
	Config *util.Config
	Client pb.CollectServiceClient
}

func NewCollectServicer(config *util.Config) (*CollectServicer, error) {
	logger := logger.NewFmtLogger()
	// Connect to the gRPC server
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.CollectService.ServerAddress, config.CollectService.GRPCServerPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not connect to gRPC server: %v", err)
	} else {
		logger.Printf("Connected to collect svc at %s:%d\n", config.CollectService.ServerAddress, config.CollectService.GRPCServerPort)
	}

	client := pb.NewCollectServiceClient(conn)
	return &CollectServicer{
		Config: config,
		Client: client,
	}, nil
}
