package controllers

import (
	"net/http"

	"github.com/caiomp87/catalog-api/internal/use-cases/product_use_case"
	"github.com/gin-gonic/gin"
)

func ListProducts(ctx *gin.Context) {
	products, err := product_use_case.ProductUseCase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := product_use_case.ProductUseCase.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func GetProductByCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	products, err := product_use_case.ProductUseCase.FindByCategoryID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func CreateProduct(ctx *gin.Context) {
	var product product_use_case.CreateProductUseCaseRequest
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := product_use_case.ProductUseCase.Create(product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}
