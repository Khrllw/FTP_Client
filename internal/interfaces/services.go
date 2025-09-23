package interfaces

type FtpService interface {
	DownloadFileToMemory(host, port, user, pass, path string) ([]byte, string, error)
	UploadFile(host, port, user, pass, targetPath, filename string, content []byte) error
}
