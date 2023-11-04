package main

import (
	"context"
	"golang-crud-gin/config"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/exporter"
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
	"go.opentelemetry.io/otel"
	// "go.elastic.co/apm/module/apmhttp"
)

// @title 	Brand Service API
// @version	1.0
// @description A Brand service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api/v1
func main() {
	// mux := http.NewServeMux()

	tp, error := exporter.TracerProvider("http://localhost:14268/api/traces")
	if error != nil {
		log.Fatal().Err(error).Msg("failed to create tracer provider")
	}
	otel.SetTracerProvider(tp)

	tr := tp.Tracer("component-main")
	_, span := tr.Start(context.Background(), "main")
	defer span.End()

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
	routes := router.NewRouter(brandController, skuController, context.Background())

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
