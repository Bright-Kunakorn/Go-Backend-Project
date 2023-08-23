package router

import (
	brandController "golang-crud-gin/pkg/brand/controller"
	skuController "golang-crud-gin/pkg/sku/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(brandController *brandController.BrandController, skuController *skuController.SkuController) *gin.Engine {
	router := gin.Default()
	// Add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	
	baseRouter := router.Group("/api/v1")

	brandsRouter := baseRouter.Group("/brand")
	brandsRouter.GET("", brandController.FindAll)
	brandsRouter.GET("/:brandId", brandController.FindById)

	skuRouter := baseRouter.Group("/sku")
	skuRouter.GET("", skuController.FindAll)
	skuRouter.GET("/:skuId", skuController.FindById)

	return router
}
