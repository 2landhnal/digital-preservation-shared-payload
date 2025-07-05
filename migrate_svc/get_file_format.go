package migrate_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (migrateServiceClient *MigrateServiceClient) GetFileFormat(ctx context.Context, req *pb.GetFileFormatRequest) (*pb.GetFileFormatResponse, error) {
	resp, err := migrateServiceClient.Client.GetFileFormat(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call GetFileFormat GAPI: %v", err)
	}

	return resp, nil
}
