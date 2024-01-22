package controller

import (
	"github.com/eduardonakaidev/go-restaurant/models"
	"github.com/eduardonakaidev/go-restaurant/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
func CreateProduct(c *gin.Context) {
	// defini o product
	var product models.Product
	// parse do bodyjson para o product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// valida todos os campos 
	if err := models.ValidateProduct(&product); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	// executa o serivce CreateProduct passando o product
	createdProduct, err := service.CreateProduct(product)
	if err != nil {
		// retorna erro caso tenha
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	// retorna o product criado e sucesso
	c.JSON(200, createdProduct)
}

func GetProductById(c *gin.Context){
	// pega o parametro id passado na url da requisição
	id := c.Param("id")
	
	//log.Println(id)
	// transforma o id em um type uuid
	iduuid,_ := uuid.Parse(id)
	//log.Println(iduuid)
	// executa o service GetByIdProduct
	productsearched, err := service.GetByIDProduct(iduuid)
	//log.Println(productsearched)
	// verifica se não retornou erro na execução do service
	if err != nil {
		// retorna o erro caso exista
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}
	// retorna o resultado caso correto
	c.JSON(200,productsearched)
}