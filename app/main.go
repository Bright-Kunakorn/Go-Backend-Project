package main

import (
	"log"

	"github.com/gin-gonic/gin"
	// _ "github.com/shall-we-go/go-gin-swagger-example/docs"
	brand "github.com/shall-we-go/go-gin-swagger-example/pkgs/brand/handlers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	
)

// @title Brand API
// @version 1.0
// @description.markdown
// @termsOfService http://somewhere.com/

// @contact.name API Support
// @contact.url http://somewhere.com/support
// @contact.email support@somewhere.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()


	r.GET("/healthcheck", brand.HealthCheckHandler)

	v1 := r.Group("/api/v1")
	{
		brands := v1.Group("/brand")
		{
			
			brandsHandler := brand.BrandsHandler{}

			// Register the GetListBrands handler
			brands.GET("", func(c *gin.Context) {
				tr := otel.Tracer("component-route")
				ctx, span := tr.Start(c.Request.Context(), "/brand")
				brandsHandler.GetListBrands(ctx, c)
		
				span.SetAttributes(
					attribute.Int("http.status_code", c.Writer.Status()),
					attribute.String("http.method", c.Request.Method),
					attribute.String("path", c.Request.URL.Path),
				)
				defer span.End()
			})

		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal((r.Run(":8080")))
}
