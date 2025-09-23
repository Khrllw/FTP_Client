package usecases

import (
	"encoding/base64"
	"ftp_client/internal/domain/models"
	"ftp_client/pkg/errors"
)

func (uc *FtpUsecase) SendFileWithAuth(req models.SendFileRequest) *errors.AppError {
	// Декодируем base64 → []byte
	data, err := base64.StdEncoding.DecodeString(req.ContentBase64)
	if err != nil {
		return errors.NewAppError(400, "Invalid base64 file content", err, false)
	}

	// Загружаем через сервис
	err2 := uc.service.UploadFile(req.Host, req.Port, req.Username, req.Password, req.TargetPath, req.Filename, data)
	if err2 != nil {
		return errors.NewAppError(500, "Error uploading file", err, false)
	}

	return nil
}
