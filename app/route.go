package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a app) Routes() {
	v1 := a.route.Group("/api/v1")

	restoran := v1.Group("restoran")

	restoran.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	})
}
