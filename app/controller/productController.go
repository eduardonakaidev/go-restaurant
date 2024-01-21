package controller

import (
	"github.com/eduardonakaidev/go-restaurant/models"
	"github.com/eduardonakaidev/go-restaurant/service"
	"github.com/gin-gonic/gin"
)
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := models.ValidateProduct(&product); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	createdProduct, err := service.CreateProduct(product)
	if err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, createdProduct)
}

func GetProducts(c *gin.Context) {
	products, err := service.GetProducts()
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, products)
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	product, err := service.GetProductById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	productindb, err := service.GetProductById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	productindb.Name = product.Name
	productindb.Description = product.Description
	productindb.Stock =product.Stock
	productindb.PriceInCents = product.PriceInCents
	if err := models.ValidateProduct(&productindb); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	updatedProduct, err := service.UpdateProduct(productindb)
	if err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"product" : updatedProduct,
		"message" : "Inserido com sucesso",
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	product, err := service.GetProductById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if err := service.DeleteProduct(product); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}