package handlers

import (
	"ftp_service/internal/config"
	"ftp_service/internal/interfaces"
	"ftp_service/internal/middleware/logging"
	"ftp_service/internal/middleware/swagger"
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

	// Подключение
	connectGroup := baseRouter.Group("/get")
	connectGroup.POST("/", h.GetFileAnonymous)    // Получить файл по FTP через соединение без авторизации
	connectGroup.POST("/auth", h.GetFileWithAuth) // Получить файл по FTP через соединение с авторизацией

	return r
}
