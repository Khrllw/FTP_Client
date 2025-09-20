package usecases

import (
	"encoding/base64"
	"ftp_service/internal/domain/models"
	"ftp_service/internal/interfaces"
	"ftp_service/pkg/errors"
)

type FTPUseCase interface {
	GetFileWithAuth(req models.GetFileWithAuthRequest) (*models.FileResponse, *errors.AppError)
	GetFileAnonymous(req models.GetFileAnonymousRequest) (*models.FileResponse, *errors.AppError)
}

type ftpUseCase struct {
	service interfaces.FtpService
}

func NewFTPUseCase(service interfaces.FtpService) FTPUseCase {
	return &ftpUseCase{service: service}
}

func (uc *ftpUseCase) GetFileWithAuth(req models.GetFileWithAuthRequest) (*models.FileResponse, *errors.AppError) {
	data, filename, err := uc.service.DownloadFileToMemory(req.Host, req.Port, req.Username, req.Password, req.FilePath)
	if err != nil {
		return nil, errors.NewAppError(500, "", err, true)
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return &models.FileResponse{
		Filename:      filename,
		ContentBase64: encoded,
	}, nil
}

func (uc *ftpUseCase) GetFileAnonymous(req models.GetFileAnonymousRequest) (*models.FileResponse, *errors.AppError) {
	data, filename, err := uc.service.DownloadFileToMemory(req.Host, req.Port, "anonymous", "anonymous", req.FilePath)
	if err != nil {
		return nil, errors.NewAppError(500, "", err, true)
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return &models.FileResponse{
		Filename:      filename,
		ContentBase64: encoded,
	}, nil
}
