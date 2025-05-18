package sharedpayload

// Payload Verify Email
type PayloadVerifyEmail struct {
	Email string `json:"email"`
}

// Payload step upload file
type PayloadStepUploadFile struct {
	Object_id      string `json:"object_id"`
	Uploader_email string `json:"uploader_email"`
	Folder_id      string `json:"folder_id"`
	Original_name  string `json:"original_name"`
	Storage_path   string `json:"storage_path"`
}
