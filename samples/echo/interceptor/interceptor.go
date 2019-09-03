package interceptor

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(
		func(username string, password string, context echo.Context) (bool, error) {
			if username == "admin" && password == "password" {
				return true, nil
			}

			return false, nil
		})
}
