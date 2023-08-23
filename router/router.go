package router

import (
	"context"
	brandController "golang-crud-gin/pkg/brand/controller"
	skuController "golang-crud-gin/pkg/sku/controller"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func NewRouter(brandController *brandController.BrandController, skuController *skuController.SkuController, ctx context.Context) *gin.Engine {
	tr := otel.Tracer("component-router")
	_, span := tr.Start(ctx, "Router")
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

	span.SetAttributes(
		attribute.String("StartTime", time.Now().String()),
	)
	defer span.End()

	return router
}
