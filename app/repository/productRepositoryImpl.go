package repository

import (
	"log"
	"time"
	"github.com/eduardonakaidev/models"
	"github.com/google/uuid"
)


func  CreateProduct(product models.Product) (models.Product, error) {
	product.CreatedAt = time.Now()
	product.ID, _ = uuid.NewRandom()
	err := models.Insert(product)
	if err != nil {
		log.Fatal("Erro: Erro ao inserir no banco de dados!")
	}
	createdProduct := product
	return createdProduct, nil
}
