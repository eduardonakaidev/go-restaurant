package service

import (
	"github.com/eduardonakaidev/go-restaurant/app/connection"
	"github.com/eduardonakaidev/go-restaurant/app/models"
	"github.com/eduardonakaidev/go-restaurant/app/utils"
)

func CreateProduct(product models.Product) (models.Product, error) {
	product.ID = utils.NewUUID()
	if err := connection.DB.Create(&product).Error; err != nil {
		return product, err
	}
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
	if err := connection.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}
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