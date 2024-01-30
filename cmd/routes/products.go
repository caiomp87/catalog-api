package routes

import (
	"github.com/caiomp87/catalog-api/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func AddProductsRoutes(route *gin.Engine) {
	productsRoute := route.Group("products")
	{
		productsRoute.GET("/", controllers.ListProducts)
		productsRoute.GET("/:id", controllers.GetProduct)
		productsRoute.GET("/categories/:id", controllers.GetProductByCategory)
		productsRoute.POST("/", controllers.CreateProduct)
	}
}
