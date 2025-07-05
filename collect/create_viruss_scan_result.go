package collect_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (client *CollectServicerClient) CreateVirussScanResult(ctx context.Context, req *pb.CreateVirussScanResultRequest) (*pb.CreateVirussScanResultResponse, error) {
	resp, err := client.Client.CreateVirussScanResult(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call CreateVirussScanResult GAPI: %v", err)
	}

	return resp, nil
}
