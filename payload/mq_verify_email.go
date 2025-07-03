package payload

import (
	"github.com/google/uuid"
)

// Payload Verify Email
type PayloadVerifyEmail struct {
	Email string `json:"email"`
}

type PipelineStep struct {
	Topic  string   `json:"topic" binding:"required"`
	Groups []string `json:"groups" binding:"required"`
}

type UpLoadFilePayload struct {
	ObjectId      uuid.UUID      `json:"object_id" binding:"required"`
	UploaderEmail string         `json:"uploader_email" binding:"required,email"`
	FolderId      int64          `json:"folder_id" binding:"required"`
	OriginalName  string         `json:"original_name" binding:"required"`
	StoragePath   string         `json:"storage_path" binding:"required"`
	Pipeline      []PipelineStep `json:"pipeline" binding:"required"` // Pipeline steps to process the file
}

func (payload *UpLoadFilePayload) NumberOfStep() int {
	return len(payload.Pipeline)
}

func (payload *UpLoadFilePayload) PopFirstPipelineStep() *PipelineStep {
	if len(payload.Pipeline) == 0 {
		return nil
	}
	step := payload.Pipeline[0]
	payload.Pipeline = payload.Pipeline[1:]
	return &step
}
