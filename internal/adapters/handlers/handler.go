package handlers

import (
	"fmt"
	"ftp_service/internal/domain/models"
	"github.com/gin-gonic/gin"
)

// GetFileWithAuth скачивает файл с FTP по логину/паролю
// @Summary Скачать файл с FTP (авторизация)
// @Description Подключается к FTP-серверу с использованием логина и пароля, и скачивает указанный файл
// @Tags FTP
// @Accept json
// @Produce json
// @Param input body models.GetFileWithAuthRequest true "Параметры подключения и файл"
// @Success 200 {object} models.FileResponse "Файл успешно загружен"
// @Failure 400 {object} swagger.IncorrectFormatError "Неверный формат запроса"
// @Failure 404 {object} swagger.NotFoundError "Файл не найден"
// @Failure 500 {object} swagger.InternalServerError "Внутренняя ошибка сервера"
// @Router /api/v1/ftp/get/auth [post]
func (h *Handler) GetFileWithAuth(c *gin.Context) {
	var req models.GetFileWithAuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err)
		return
	}

	resp, eerr := h.usecase.GetFileWithAuth(req)
	if eerr != nil {
		h.ErrorResponse(c, eerr, eerr.Code, eerr.Message, true)
		return
	}

	h.ResultResponse(c, fmt.Sprintf("File got successfull"), Object, resp)
}

// GetFileAnonymous скачивает файл с FTP анонимно
// @Summary Скачать файл с FTP (анонимный доступ)
// @Description Подключается к FTP-серверу как anonymous и скачивает указанный файл
// @Tags FTP
// @Accept json
// @Produce json
// @Param input body models.GetFileAnonymousRequest true "Адрес сервера и путь к файлу"
// @Success 200 {object} models.FileResponse "Файл успешно загружен"
// @Failure 400 {object} swagger.IncorrectFormatError "Неверный формат запроса"
// @Failure 404 {object} swagger.NotFoundError "Файл не найден"
// @Failure 500 {object} swagger.InternalServerError "Внутренняя ошибка сервера"
// @Router /api/v1/ftp/get/anon [post]
func (h *Handler) GetFileAnonymous(c *gin.Context) {
	var req models.GetFileAnonymousRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.BadRequest(c, err)
		return
	}

	resp, eerr := h.usecase.GetFileAnonymous(req)
	if eerr != nil {
		h.ErrorResponse(c, eerr, eerr.Code, eerr.Message, true)
		return
	}

	h.ResultResponse(c, fmt.Sprintf("File got successfull"), Object, resp)
}
