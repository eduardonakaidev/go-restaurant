package main

import (
	"github.com/eduardonakaidev/go-restaurant/app/controller"
	"github.com/gofiber/fiber/v2"
)




func main() {
	app := fiber.New()
	controller.HelloController(app)	
	app.Listen(":5050")
}
