package interfaces

import (
	"ftp_client/internal/domain/models"
	"ftp_client/pkg/errors"
)

type Usecases interface {
	FtpUsecase
}

type FtpUsecase interface {
	GetFileWithAuth(req models.GetFileWithAuthRequest) (*models.FileResponse, *errors.AppError)
	GetFileAnonymous(req models.GetFileAnonymousRequest) (*models.FileResponse, *errors.AppError)
	SendFileWithAuth(req models.SendFileRequest) *errors.AppError
}
