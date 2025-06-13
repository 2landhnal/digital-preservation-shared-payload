package core_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (coreServiceClient *CoreServiceClient) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	resp, err := coreServiceClient.Client.UploadFile(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call UploadFile GAPI: %v", err)
	}

	return resp, nil
}
