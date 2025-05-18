package payload

import (
	"github.com/google/uuid"
)

// Payload Verify Email
type PayloadVerifyEmail struct {
	Email string `json:"email"`
}

// Payload step upload file
type PayloadStepUploadFile struct {
	ObjectId      uuid.UUID `json:"object_id" binding:"required"`
	UploaderEmail string    `json:"uploader_email" binding:"required,email"`
	FolderId      int64     `json:"folder_id" binding:"required"`
	OriginalName  string    `json:"original_name" binding:"required"`
	StoragePath   string    `json:"storage_path" binding:"required"`
}
