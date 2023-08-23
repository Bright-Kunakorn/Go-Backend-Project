package main

import (
	"golang-crud-gin/config"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	brandController "golang-crud-gin/pkg/brand/controller"
	brandModel "golang-crud-gin/pkg/brand/model"
	brandRepository "golang-crud-gin/pkg/brand/repository"
	brandService "golang-crud-gin/pkg/brand/service"
	skuController "golang-crud-gin/pkg/sku/controller"
	skuModel "golang-crud-gin/pkg/sku/model"
	skuRepository "golang-crud-gin/pkg/sku/repository"
	skuService "golang-crud-gin/pkg/sku/service"
	"golang-crud-gin/router"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Brand Service API
// @version	1.0
// @description A Brand service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api/v1
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("brand").AutoMigrate(&brandModel.Brand{})
	db.Table("sku").AutoMigrate(&skuModel.Sku{})

	// Repository
	brandRepository := brandRepository.NewBrandREpositoryImpl(db)
	skuRepository := skuRepository.NewSkuREpositoryImpl(db)

	// Service
	brandService := brandService.NewBrandServiceImpl(brandRepository, validate)
	skuService := skuService.NewSkuServiceImpl(skuRepository, validate)

	// Controller
	brandController := brandController.NewBrandController(brandService)
	skuController := skuController.NewSkuController(skuService)

	// Router
	routes := router.NewRouter(brandController, skuController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

// package main

// import (
// 	"net/http"

// 	brandController "golang-crud-gin/pkg/brand/controller"
// 	brandModel "golang-crud-gin/pkg/brand/model"
// 	brandRepository "golang-crud-gin/pkg/brand/repository"
// 	brandRouter "golang-crud-gin/pkg/brand/router"
// 	brandService "golang-crud-gin/pkg/brand/service"

// 	skuController "golang-crud-gin/pkg/sku/controller"
// 	skuModel "golang-crud-gin/pkg/sku/model"
// 	skuRepository "golang-crud-gin/pkg/sku/repository"
// 	skuRouter "golang-crud-gin/pkg/sku/router"
// 	skuService "golang-crud-gin/pkg/sku/service"

// 	"golang-crud-gin/config"
// 	_ "golang-crud-gin/docs"
// 	"golang-crud-gin/helper"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"github.com/rs/zerolog/log"
// )

// // @title 	Inventory Service API
// // @version	1.0
// // @description A Inventory service API in Go using Gin framework

// // @host 	localhost:8888
// // @BasePath /api/v1
// func main() {
// 	log.Info().Msg("Started Server!")

// 	// Database
// 	db := config.DatabaseConnection()
// 	validate := validator.New()

// 	// Brand-related
// 	db.Table("brand").AutoMigrate(&brandModel.Brand{})
// 	brandRepository := brandRepository.NewBrandREpositoryImpl(db)
// 	brandService := brandService.NewBrandServiceImpl(brandRepository, validate)
// 	brandController := brandController.NewBrandController(brandService)

// 	// SKU-related
// 	db.Table("sku").AutoMigrate(&skuModel.Backendposdatasku{})
// 	skuRepository := skuRepository.NewSkuREpositoryImpl(db)
// 	skuService := skuService.NewSkuServiceImpl(skuRepository, validate)
// 	skuController := skuController.NewSkuController(skuService)

// 	// Create brand router using the NewRouter function from brand router package
// 	brandRouter := brandRouter.NewRouter(brandController)

// 	// Create SKU router using the NewRouter function from SKU router package
// 	skuRouter := skuRouter.NewRouter(skuController)

// 	mainRouter := gin.Default()

// 	// Map brand routes
// 	mainRouter.GET("/brand/*any", func(c *gin.Context) {
// 		brandRouter.ServeHTTP(c.Writer, c.Request)
// 	})

// 	// Map SKU routes
// 	mainRouter.GET("/sku/*any", func(c *gin.Context) {
// 		skuRouter.ServeHTTP(c.Writer, c.Request)
// 	})

// 	server := &http.Server{
// 		Addr:    ":8888",
// 		Handler: mainRouter,
// 	}

// 	err := server.ListenAndServe()
// 	helper.ErrorPanic(err)
// }
