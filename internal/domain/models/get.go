package models

type GetFileAnonymousRequest struct {
	Host     string `json:"host" example:"192.168.1.1"`
	Port     string `json:"port" example:"22"`
	FilePath string `json:"file_path" example:"/data/test.txt"`
}

type GetFileWithAuthRequest struct {
	Host     string `json:"host" example:"192.168.1.1"`
	Port     string `json:"port" example:"22"`
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password"`
	FilePath string `json:"file_path" example:"/data/test.txt"`
}

type FileResponse struct {
	Filename      string `json:"filename" example:"test.txt"`
	ContentBase64 string `json:"content_base64" example:"MQ0KMg0KMw0KdGVzdCBzdWNjZXNzISE="`
}
