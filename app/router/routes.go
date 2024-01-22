package router

import (
	"github.com/gin-gonic/gin"
	"github.com/eduardonakaidev/go-restaurant/controller"
)

func HandleRequest() {
	r := gin.Default()
	api := r.Group("/api")
	{
		product := api.Group("/product")
		{
			product.POST("/", controller.CreateProduct)
			product.GET("/:id",controller.GetProductById)
		}
	}
	r.Run(":5050")
}