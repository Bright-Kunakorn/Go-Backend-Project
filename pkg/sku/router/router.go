package router

import (
	"golang-crud-gin/pkg/sku/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(skuController *controller.SkuController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api/v1")
	skuRouter := baseRouter.Group("/sku")
	skuRouter.GET("", skuController.FindAll)
	skuRouter.GET("/:skuId", skuController.FindById)

	return router
}
