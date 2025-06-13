package collect_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (client *CollectServicer) CreateExtractMetadataResult(ctx context.Context, req *pb.CreateExtractMetadataResultRequest) (*pb.CreateExtractMetadataResultResponse, error) {
	resp, err := client.Client.CreateExtractMetadataResult(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call CreateExtractMetadataResult GAPI: %v", err)
	}

	return resp, nil
}
