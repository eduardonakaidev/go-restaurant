package main

import (
	"github.com/eduardonakaidev/configs"
	"github.com/eduardonakaidev/controller"
	"github.com/eduardonakaidev/middleware"
	"github.com/gofiber/fiber/v2"
)




func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	conn, _:= configs.OpenConnection()
	query := `
    CREATE TABLE IF NOT EXISTS product (
		id UUID PRIMARY KEY,
		name VARCHAR(255),
		description VARCHAR(255),
		priceincents INT,
		stock INT
	);`	
	
	conn.Query(query)
	
	app := fiber.New()
	app.Use(middleware.SetContentType("application/json"))
	controller.ProductController(app)
	app.Listen(":5050")
}
