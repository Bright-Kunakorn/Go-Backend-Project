package router

import (
	"golang-crud-gin/pkg/brand/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(brandsController *controller.BrandsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api/v1")
	brandsRouter := baseRouter.Group("/brands")
	brandsRouter.GET("", brandsController.FindAll)
	brandsRouter.GET("/:brandId", brandsController.FindById)

	return router
}
