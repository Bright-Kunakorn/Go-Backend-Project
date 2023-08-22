package main

import (
	"golang-crud-gin/config"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/controller"
	"golang-crud-gin/pkg/brand/model"
	"golang-crud-gin/pkg/brand/repository"
	"golang-crud-gin/pkg/brand/router"
	"golang-crud-gin/pkg/brand/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Brand Service API
// @version	1.0
// @description A Brand service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	log.Info().Msg("Started Server!")
	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("brand").AutoMigrate(&model.Brands{})

	// Repository
	brandsRepository := repository.NewBrandsREpositoryImpl(db)

	// Service
	brandsService := service.NewBrandsServiceImpl(brandsRepository, validate)

	// Controller
	brandsController := controller.NewBrandsController(brandsService)

	// Router
	routes := router.NewRouter(brandsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
