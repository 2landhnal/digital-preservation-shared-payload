package collect_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (client *CollectServicer) UpdateTmpObject(ctx context.Context, req *pb.UpdateTmpObjectRequest) (*pb.UpdateTmpObjectResponse, error) {
	resp, err := client.Client.UpdateTmpObject(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call UpdateTmpObject GAPI: %v", err)
	}

	return resp, nil
}
