package interfaces

import (
	"ftp_service/internal/domain/models"
	"ftp_service/pkg/errors"
)

type Usecases interface {
	FtpUsecase
}

type FtpUsecase interface {
	GetFileWithAuth(req models.GetFileWithAuthRequest) (*models.FileResponse, *errors.AppError)
	GetFileAnonymous(req models.GetFileAnonymousRequest) (*models.FileResponse, *errors.AppError)
}
