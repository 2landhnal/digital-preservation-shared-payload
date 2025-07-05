package core_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (coreServiceClient *CoreServiceClient) GetPresignedUrl(ctx context.Context, req *pb.GetPresignedUrlRequest) (*pb.GetPresignedUrlResponse, error) {
	resp, err := coreServiceClient.Client.GetPresignedUrl(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call GetPresignedUrl GAPI: %v", err)
	}

	return resp, nil
}
