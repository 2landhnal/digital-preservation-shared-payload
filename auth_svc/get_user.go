package auth_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/logger"
	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (authServiceClient *AuthServiceClient) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Call the Identify RPC
	logger := logger.NewFmtLogger()
	logger.Println("Getting user with email:", req.Email)
	resp, err := authServiceClient.Client.GetUser(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call GetUser GAPI: %v", err)
	}

	return resp, nil
}
