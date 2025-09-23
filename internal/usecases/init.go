package usecases

import (
	"ftp_client/internal/interfaces"
	_ "ftp_client/pkg/errors"
)

type UseCases struct {
	interfaces.FtpUsecase
}

func NewUsecases(s interfaces.FtpService) interfaces.Usecases {
	return &UseCases{
		NewFTPUsecase(s),
	}
}
