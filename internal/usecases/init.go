package usecases

import (
	"ftp_service/internal/interfaces"
	_ "ftp_service/pkg/errors"
)

type UseCases struct {
	interfaces.FtpUsecase
}

func NewUsecases(s interfaces.FtpService) interfaces.Usecases {
	return &UseCases{
		NewFTPUseCase(s),
	}
}
