package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetEventsServiceHealth - Проверка работоспособности микросервиса событий
func (c *Container) GetEventsServiceHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]bool{"status": true})
}
