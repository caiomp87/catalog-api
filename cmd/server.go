package main

import (
	"database/sql"
	"fmt"

	"github.com/caiomp87/catalog-api/cmd/routes"
	"github.com/caiomp87/catalog-api/internal/env"
	"github.com/caiomp87/catalog-api/internal/repositories/category_repository"
	"github.com/caiomp87/catalog-api/internal/repositories/product_repository"
	"github.com/caiomp87/catalog-api/internal/use-cases/category_use_case"
	"github.com/caiomp87/catalog-api/internal/use-cases/product_use_case"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func init() {
	logger := zap.Must(zap.NewDevelopment())

	appEnv := env.GetEnv("APP_ENV").FallbackString("development")

	fmt.Println(appEnv)

	if appEnv == "production" {
		logger = zap.Must(zap.NewProduction())
	}

	zap.ReplaceGlobals(logger)
}

func main() {
	driver := env.GetEnv("DATABASE_DRIVE").FallbackString("postgres")
	dbUrl := env.GetEnv("DATABASE_URL").String()

	conn, err := sql.Open(driver, dbUrl)
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

	port := env.GetEnv("API_PORT").FallbackString("3333")

	zap.L().Info("[x] API is running", zap.String("port", port))

	if err := server.Run(":" + port); err != nil {
		zap.L().Fatal("failed to serve", zap.Error(err))
	}
}
