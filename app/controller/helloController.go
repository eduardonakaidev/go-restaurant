package controller

import (
	"github.com/eduardonakaidev/go-restaurant/handler"
	"github.com/gofiber/fiber/v2"
)


func HelloController(c *fiber.App) {
	c.Get("/", handler.HelloHandler)
}