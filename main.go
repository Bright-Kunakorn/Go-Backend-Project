package main

import (
	"golang-crud-gin/config"
	"golang-crud-gin/controller"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/router"
	"golang-crud-gin/service"
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

	db.Table("brands").AutoMigrate(&model.Brands{})

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
