package handlers

import (
	_ "ftp_client/docs"
	"ftp_client/internal/config"
	"ftp_client/internal/interfaces"
	"ftp_client/internal/middleware/logging"
	"ftp_client/internal/middleware/swagger"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	"net/http"
)

type Handler struct {
	logger  *logging.Logger
	usecase interfaces.Usecases
	service interfaces.FtpService
}

// NewHandler создает новый экземпляр Handler со всеми зависимостями
func NewHandler(usecase interfaces.Usecases, parentLogger *logging.Logger, service interfaces.FtpService) *Handler {
	handlerLogger := parentLogger.WithPrefix("HANDLER")
	handlerLogger.Info("Handler initialized",
		"component", "GENERAL",
	)
	return &Handler{
		logger:  handlerLogger,
		usecase: usecase,
		service: service,
	}
}

// ProvideRouter создает и настраивает маршруты
func ProvideRouter(h *Handler, cfg *config.Config, swagCfg *swagger.Config) http.Handler {
	gin.SetMode(cfg.App.GinMode)
	r := gin.Default()

	// Swagger-роутер
	swagger.Setup(r, swagCfg)

	// LoggerMiddleware
	r.Use(LoggingMiddleware(h.logger))

	// Общая группа для API
	baseRouter := r.Group("/api/v1")

	// Получение
	getGroup := baseRouter.Group("/get")
	getGroup.POST("/anon", h.GetFileAnonymous) // Получить файл по FTP через соединение без авторизации
	getGroup.POST("/auth", h.GetFileWithAuth)  // Получить файл по FTP через соединение с авторизацией

	// Отправка
	sendGroup := baseRouter.Group("/send")
	sendGroup.POST("/anon", h.SendFileAnonymous) // Получить файл по FTP через соединение без авторизации
	sendGroup.POST("/auth", h.SendFileWithAuth)  // Получить файл по FTP через соединение с авторизацией

	return r
}
