package models

type SendFileRequest struct {
	Host          string `json:"host" binding:"required" example:"192.168.30.92"`
	Port          string `json:"port" binding:"required" example:"21"`
	Username      string `json:"username,omitempty" example:"tester"`
	Password      string `json:"password,omitempty" example:"password"`
	TargetPath    string `json:"target_path" binding:"required" example:"upload/"`
	Filename      string `json:"filename" binding:"required" example:"example.txt"`
	ContentBase64 string `json:"content_base64" binding:"required" example:"SGVsbG8sIEZUUCE="`
}
