package services

import (
	"bytes"
	"fmt"
	"ftp_client/internal/interfaces"
	"ftp_client/internal/middleware/logging"
	"github.com/jlaffaye/ftp"
	"io"
	"path/filepath"
	"strings"
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

func (s *FtpService) UploadFile(host, port, user, pass, targetPath, filename string, content []byte) error {
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return err
	}
	defer conn.Quit()

	// Авторизация
	if err := conn.Login(user, pass); err != nil {
		return err
	}

	// Переход в нужную директорию
	if targetPath != "" && targetPath != "." {
		if err := conn.MakeDir(targetPath); err != nil && !ftpErrIsFileExists(err) {
			s.logger.Error("MakeDir failed: %v", err)
			return err
		}
		if err := conn.ChangeDir(targetPath); err != nil {
			s.logger.Error("ChangeDir failed: %v", err)
			return err
		}
	}

	// Загрузка
	reader := bytes.NewReader(content)
	if err := conn.Stor(filename, reader); err != nil {
		s.logger.Error("File upload failed: %v", err)
		return err
	}

	return nil
}

// ftpErrIsFileExists проверяет, существует ли директория (для создания папки)
func ftpErrIsFileExists(err error) bool {
	// Простой хак для популярных FTP серверов. Можно кастомизировать под твой.
	return err != nil && (strings.Contains(err.Error(), "file exists") || strings.Contains(err.Error(), "550"))
}

func getFilenameFromPath(path string) string {
	return filepath.Base(path)
}
