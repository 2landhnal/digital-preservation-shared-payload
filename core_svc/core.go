package core_svc

import (
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/logger"
	"github.com/2landhnal/digital-preservation-shared-payload/pb"
	"github.com/2landhnal/digital-preservation-shared-payload/util"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CoreServiceClient struct {
	Config *util.Config
	Client pb.CoreServiceClient
}

func NewCoreServiceClient(config *util.Config) (*CoreServiceClient, error) {
	logger := logger.NewFmtLogger()
	// Connect to the gRPC server
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", config.CoreService.ServerAddress, config.CoreService.GRPCServerPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, fmt.Errorf("could not connect to gRPC server: %v", err)
	} else {
		logger.Printf("Connected to gRPC server at %s:%d\n", config.CoreService.ServerAddress, config.CoreService.GRPCServerPort)
	}

	client := pb.NewCoreServiceClient(conn)
	return &CoreServiceClient{
		Config: config,
		Client: client,
	}, nil
}
