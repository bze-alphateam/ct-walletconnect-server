package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func registerRoutes(srv *echo.Echo, ctrl *ControllerFactory) {
	srv.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{}{})
	})
}
