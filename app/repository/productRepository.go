package repository

import "github.com/eduardonakaidev/models"

type ProductRepository interface {
	CreateProduct(product models.Product) (models.Product, error)
}
