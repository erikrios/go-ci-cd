package middleware

import (
	"github.com/erikrios/go-clean-arhictecture/constants"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		SigningKey: []byte(constants.JWT_SECRET),
	}

	return middleware.JWTWithConfig(config)
}
