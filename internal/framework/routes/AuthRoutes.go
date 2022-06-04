package routes

import (
	"go-hospital-server/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewAuthRoutes(e *echo.Group, acon *controller.AuthController, middleware ...echo.MiddlewareFunc) {
	e.POST("/refresh_token", acon.RefreshToken)
	e.POST("/login", acon.Login)
}