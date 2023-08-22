package router

import (
	"golang-crud-gin/pkg/brand/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(brandController *controller.BrandController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api/v1")
	brandsRouter := baseRouter.Group("/brand")
	brandsRouter.GET("", brandController.FindAll)
	brandsRouter.GET("/:brandId", brandController.FindById)

	return router
}
