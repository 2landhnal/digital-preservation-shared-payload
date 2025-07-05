package migrate_svc

import (
	"context"
	"fmt"

	"github.com/2landhnal/digital-preservation-shared-payload/pb"
)

func (migrateServiceClient *MigrateServiceClient) CheckFileFormatAcceptable(ctx context.Context, req *pb.CheckFileFormatAcceptableRequest) (*pb.CheckFileFormatAcceptableResponse, error) {
	resp, err := migrateServiceClient.Client.CheckFileFormatAcceptable(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("could not call CheckFileFormatAcceptable GAPI: %v", err)
	}

	return resp, nil
}
