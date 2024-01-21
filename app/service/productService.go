package service

import (
	"time"

	"github.com/eduardonakaidev/go-restaurant/connection"
	"github.com/eduardonakaidev/go-restaurant/models"
	"github.com/eduardonakaidev/go-restaurant/utils"
)

func CreateProduct(product models.Product) (models.Product, error) {
	product.ID = utils.NewUUID()
	product.CreatedAt = time.Now()
	if err := connection.DB.QueryRow()
	return product, nil
}

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	if err := connection.DB.Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product
	if err := connection.DB.QueryRow()
	return product, nil
}

func UpdateProduct(product models.Product) (models.Product, error) {
	if err := connection.DB.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func DeleteProduct(product models.Product) error {
	if err := connection.DB.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}