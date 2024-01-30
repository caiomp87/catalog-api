package routes

import (
	"github.com/caiomp87/catalog-api/cmd/controllers"
	"github.com/gin-gonic/gin"
)

func AddCategoriesRoutes(route *gin.Engine) {
	categoriesRoute := route.Group("categories")
	{
		categoriesRoute.GET("/", controllers.ListCategories)
		categoriesRoute.GET("/:id", controllers.GetCategory)
		categoriesRoute.POST("/", controllers.CreateCategory)
	}
}
