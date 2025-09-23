package swagger

import "ftp_client/internal/domain/models"

type SendResponse struct {
	Status string `json:"status" example:"ok"`

	Message string `json:"message" example:"MSG"`
	Type    string `json:"type" example:"Empty"`
}

type GetResponse struct {
	Status  string              `json:"status" example:"ok"`
	Message string              `json:"message" example:"Successfully connected"`
	Type    string              `json:"type" example:"object"`
	Data    models.FileResponse `json:"data"`
}
