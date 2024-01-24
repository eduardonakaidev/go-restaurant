package controller

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// valida todos os campos
	if err := models.ValidateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// executa o serivce CreateProduct passando o product
	createdProduct, err := service.CreateProduct(product)
	if err != nil {
		// retorna erro caso tenha
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// retorna o product criado e sucesso
	c.JSON(http.StatusOK, createdProduct)
}

func GetProductById(c *gin.Context) {
	// pega o parametro id passado na url da requisição
	id := c.Param("id")

	//log.Println(id)
	// transforma o id em um type uuid
	iduuid, _ := uuid.Parse(id)
	//log.Println(iduuid)
	// executa o service GetByIdProduct
	productsearched, err := service.GetByIDProduct(iduuid)
	//log.Println(productsearched)
	// verifica se não retornou erro na execução do service
	if err != nil {
		// retorna o erro caso exista
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// retorna o resultado caso correto
	c.JSON(http.StatusOK, productsearched)
}

func ListProduct(c *gin.Context) {
	products, err := service.ListProduct()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
func Update(c *gin.Context) {
	// pega o parametro id passado na url da requisição
	id := c.Param("id")

	// transforma o id em um type uuid
	iduuid, _ := uuid.Parse(id)

	// defini o product
	var productdto models.UpdateProductDTO
	// parse do bodyjson para o product
	if err := c.ShouldBindJSON(&productdto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product, err := service.UpdateProduct(iduuid,productdto)
	// verifica se não retornou erro na execução do service
	if err != nil {
		// retorna o erro caso exista
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
func DeleteProductController(c *gin.Context) {
	// Obter o ID do produto a ser excluído dos parâmetros da rota
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de produto inválido"})
		return
	}

	// Tentar excluir o produto
	err = service.DeleteProduct(id)
	if err != nil {
		if err.Error() == "Produto não encontrado" {
			// Se o erro for "Produto não encontrado", retornar 404 Not Found
			c.JSON(http.StatusBadRequest, gin.H{"error": "Produto não encontrado"})
		} else {
			// Para outros erros, retornar 500 Internal Server Error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir o produto"})
		}
		return
	}

	// Produto excluído com sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Produto excluído com sucesso"})
}