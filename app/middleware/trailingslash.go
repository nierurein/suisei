package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UseTrailingSlash(e *echo.Echo) {
	e.Pre(middleware.AddTrailingSlash())
	e.Pre(middleware.RemoveTrailingSlash())
}
