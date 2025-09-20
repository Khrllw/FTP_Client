package interfaces

type FtpService interface {
	DownloadFileToMemory(host, port, user, pass, path string) ([]byte, string, error)
}
