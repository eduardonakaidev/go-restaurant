package service

import (
	"github.com/eduardonakaidev/models"
	"github.com/eduardonakaidev/repository"
)

var productRepo repository.ProductRepository
func CreateProduct(product models.Product)(models.Product,error){
	createdProduct, err := productRepo.CreateProduct(product)
	if err != nil {
		return models.Product{},err
	}
	return createdProduct,nil
}
