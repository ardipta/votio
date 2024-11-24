package routes

import (
	"github.com/ashiqur/votio/internal/api"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", api.HomePageHandler)
}
