package handler

import (
	"encoding/json"
	"log"

	"github.com/eduardonakaidev/models"
	"github.com/eduardonakaidev/service"
	"github.com/gofiber/fiber/v2"
)

func CreateProductHandler (c *fiber.Ctx)error{

		var product models.Product
		err := json.Unmarshal(c.Body(), &product)
		log.Println(err)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request body invalid"})
		}
		createdProduct, err := service.CreateProduct(product)

		if err != nil{
			log.Printf("Erro ao criar produto: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Erro ao criar produto"})
		}
		
		return c.Status(fiber.StatusCreated).JSON(createdProduct)
}