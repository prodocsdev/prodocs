package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Prodocs!")
}

func Path(c echo.Context) error {
	path := c.Param("path")
	return c.String(http.StatusOK, path)
}

func Hook(c echo.Context) error {
	var event map[string]interface{}
	if err := c.Bind(&event); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, event)
}
