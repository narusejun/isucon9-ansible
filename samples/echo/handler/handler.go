package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func SubPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello World"+username)
	}
}

func JsonPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		whatever := c.Param("whatever")

		jsonMap := map[string]string{
			"whatever": whatever,
		}

		return c.JSON(http.StatusOK, jsonMap)
	}
}

func returnJson(c echo.Context) error {
	whenever := c.Param("whatever")

	jsonMap := map[string]string{
		"whenever": whenever,
	}

	return c.JSON(http.StatusOK, jsonMap)
}
