package collect_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (client *CollectServicerClient) CreateTmpObject(ctx context.Context, req *pb.CreateTmpObjectRequest) (*pb.CreateTmpObjectResponse, error) {
	resp, err := client.Client.CreateTmpObject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call CreateTmpObject GAPI: %v", err)
	}
	return resp, nil
}
