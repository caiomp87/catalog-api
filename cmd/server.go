package main

import (
	"database/sql"
	"os"

	"github.com/caiomp87/catalog-api/cmd/routes"
	"github.com/caiomp87/catalog-api/internal/repositories/category_repository"
	"github.com/caiomp87/catalog-api/internal/repositories/product_repository"
	"github.com/caiomp87/catalog-api/internal/use-cases/category_use_case"
	"github.com/caiomp87/catalog-api/internal/use-cases/product_use_case"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	logger := zap.Must(zap.NewDevelopment())
	if os.Getenv("APP_ENV") == "production" {
		logger = zap.Must(zap.NewDevelopment())
	}

	zap.ReplaceGlobals(logger)
}

func main() {
	conn, err := sql.Open("postgres", "postgresql://caio:123456@localhost:5433/imersao?sslmode=disable")
	if err != nil {
		zap.L().Fatal("failed to connect to database", zap.Error(err))
	}

	categoryRepository := category_repository.NewCategoryRepository(conn)
	productRepository := product_repository.NewProductRepository(conn)

	category_use_case.CategoryUseCase = category_use_case.NewCategoryUseCase(categoryRepository)
	product_use_case.ProductUseCase = product_use_case.NewProductUseCase(productRepository)

	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	routes.AddCategoriesRoutes(server)
	routes.AddProductsRoutes(server)

	zap.L().Info("[x] API is running", zap.String("port", "3333"))

	if err := server.Run(":3333"); err != nil {
		zap.L().Fatal("failed to serve", zap.Error(err))
	}
}
