package controller

import (
	"github.com/eduardonakaidev/handler"
	"github.com/gofiber/fiber/v2"
)

func ProductController(c *fiber.App) {
	c.Post("/product",handler.CreateProductHandler)
}