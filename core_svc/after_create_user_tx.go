package core_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (CoreServiceClient *CoreServiceClient) AfterCreateUserTx(ctx context.Context, req *pb.AfterCreateUserTxRequest) (*pb.AfterCreateUserTxResponse, error) {
	resp, err := CoreServiceClient.Client.AfterCreateUserTx(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call AfterCreateUserTx GAPI: %v", err)
	}

	return resp, nil
}
