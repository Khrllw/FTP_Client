package usecases

import (
	"encoding/base64"
	"ftp_client/internal/domain/models"
	"ftp_client/internal/interfaces"
	"ftp_client/pkg/errors"
)

type FTPUsecase interface {
	GetFileWithAuth(req models.GetFileWithAuthRequest) (*models.FileResponse, *errors.AppError)
	GetFileAnonymous(req models.GetFileAnonymousRequest) (*models.FileResponse, *errors.AppError)
}

type FtpUsecase struct {
	service interfaces.FtpService
}

func NewFTPUsecase(service interfaces.FtpService) *FtpUsecase {
	return &FtpUsecase{service: service}
}

func (uc *FtpUsecase) GetFileWithAuth(req models.GetFileWithAuthRequest) (*models.FileResponse, *errors.AppError) {
	data, filename, err := uc.service.DownloadFileToMemory(req.Host, req.Port, req.Username, req.Password, req.FilePath)
	if err != nil {
		return nil, errors.NewAppError(500, errors.InternalServerError, err, false)
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return &models.FileResponse{
		Filename:      filename,
		ContentBase64: encoded,
	}, nil
}

func (uc *FtpUsecase) GetFileAnonymous(req models.GetFileAnonymousRequest) (*models.FileResponse, *errors.AppError) {
	data, filename, err := uc.service.DownloadFileToMemory(req.Host, req.Port, "anonymous", "anonymous", req.FilePath)
	if err != nil {
		return nil, errors.NewAppError(500, errors.InternalServerError, err, false)
	}

	encoded := base64.StdEncoding.EncodeToString(data)

	return &models.FileResponse{
		Filename:      filename,
		ContentBase64: encoded,
	}, nil
}
