package services

import (
	"fmt"
	"ftp_service/internal/interfaces"
	"ftp_service/internal/middleware/logging"
	"github.com/jlaffaye/ftp"
	"io"
	"path/filepath"
	"time"
)

type FtpService struct {
	logger *logging.Logger
}

func NewFtpService(logger *logging.Logger) interfaces.FtpService {
	return &FtpService{logger: logger}
}

func (s *FtpService) DownloadFileToMemory(host, port, user, pass, path string) ([]byte, string, error) {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return nil, "", err
	}
	defer conn.Quit()

	if err := conn.Login(user, pass); err != nil {
		return nil, "", err
	}

	resp, err := conn.Retr(path)
	if err != nil {
		return nil, "", err
	}
	defer resp.Close()

	data, err := io.ReadAll(resp)
	if err != nil {
		return nil, "", err
	}

	return data, getFilenameFromPath(path), nil
}

func getFilenameFromPath(path string) string {
	return filepath.Base(path)
}
