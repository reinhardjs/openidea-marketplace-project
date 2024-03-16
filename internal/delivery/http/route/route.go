package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/openidea-marketplace/internal/delivery/http"
)

type RouteConfig struct {
	App            *fiber.App
	UserHandler    *http.UserHandler
	AuthMiddleware fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)

	v1 := c.App.Group("/v1")
	v1.Post("/user/register", c.UserHandler.Register)
}
