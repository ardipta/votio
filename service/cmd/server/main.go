package main

import (
	"github.com/ashiqur/votio/internal/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// TODO: initial config load
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
