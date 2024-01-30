package controllers

import (
	"net/http"

	"github.com/caiomp87/catalog-api/internal/use-cases/category_use_case"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context) {
	categories, err := category_use_case.CategoryUseCase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")

	category, err := category_use_case.CategoryUseCase.FindByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func CreateCategory(ctx *gin.Context) {
	var category category_use_case.CreateCategoryUseCaseRequest
	if err := ctx.BindJSON(&category); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := category_use_case.CategoryUseCase.Create(category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(http.StatusCreated)
}
