package handlers

import (
	"ftp_client/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// SendFileWithAuth загружает файл на FTP с авторизацией
// @Summary Загрузить файл на FTP (авторизация)
// @Description Подключается к FTP-серверу с логином и паролем и загружает указанный файл
// @Tags FTP_Auth
// @Accept json
// @Produce json
// @Param input body models.SendFileRequest true "Данные подключения и файл"
// @Success 200 {object} swagger.SendResponse "Файл успешно отправлен"
// @Failure 400 {object} swagger.IncorrectFormatError "Неверный формат запроса"
// @Failure 500 {object} swagger.InternalServerError "Ошибка при загрузке файла"
// @Router /send/auth [post]
func (h *Handler) SendFileWithAuth(c *gin.Context) {
	var req models.SendFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err)
		return
	}

	eerr := h.usecase.SendFileWithAuth(req)
	if eerr != nil {
		h.ErrorResponse(c, eerr, eerr.Code, eerr.Message, true)
		return
	}

	h.ResultResponse(c, "File uploaded successfully", Empty, nil)
}

// SendFileAnonymous загружает файл на FTP без авторизации
// @Summary Загрузить файл на FTP (анонимный доступ)
// @Description Подключается к FTP-серверу как anonymous и загружает файл
// @Tags FTP_Anon
// @Accept json
// @Produce json
// @Param input body models.SendFileRequest true "Данные подключения и файл"
// @Success 200 {object} swagger.SendResponse "Файл успешно отправлен"
// @Failure 400 {object} swagger.IncorrectFormatError "Неверный формат запроса"
// @Failure 500 {object} swagger.InternalServerError "Ошибка при загрузке файла"
// @Router /send/anon [post]
func (h *Handler) SendFileAnonymous(c *gin.Context) {
	var req models.SendFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err)
		return
	}

	// Подставим anonymous
	req.Username = "anonymous"
	req.Password = "anonymous"

	err := h.usecase.SendFileWithAuth(req)
	if err != nil {
		h.ErrorResponse(c, err, err.Code, err.Message, true)
		return
	}
	h.ResultResponse(c, "File uploaded successfully", Empty, nil)
}
