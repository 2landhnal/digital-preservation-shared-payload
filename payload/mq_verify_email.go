package payload

// Payload Verify Email
type PayloadVerifyEmail struct {
	Email string `json:"email"`
}

// Payload step upload file
type PayloadStepUploadFile struct {
	ObjectId      string `json:"object_id"`
	UploaderEmail string `json:"uploader_email"`
	FolderId      string `json:"folder_id"`
	OriginalName  string `json:"original_name"`
	StoragePath   string `json:"storage_path"`
}
