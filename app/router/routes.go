package router

import (
	"github.com/gin-gonic/gin"
	"github.com/eduardonakaidev/go-restaurant/controller"
)

func HandleRequest() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/ping", controller.Ping)
		product := api.Group("/product")
		{
			product.POST("/", controller.CreateProduct)
			product.GET("/", controller.GetProducts)
			product.GET("/:id", controller.GetProductById)
			product.PUT("/:id", controller.UpdateProduct)
			product.DELETE("/:id", controller.DeleteProduct)
		}
	}
	r.Run(":5050")
}