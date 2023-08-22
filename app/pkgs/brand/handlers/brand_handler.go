package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shall-we-go/go-gin-swagger-example/pkgs/brand/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type BrandsHandler struct {
}

// GetBrand godoc
// @summary Get Brand
// @description  Get brand
// @tags brand
// @security ApiKeyAuth
// @id GetBrand
// @accept json
// @produce json
// @param id path int true "id of brand to be gotten"
// @response 200 {object} model.brandModel "OK"
// @response 400 {object} model.brandModel "Bad Request"
// @response 401 {object} model.brandModel "Unauthorized"
// @response 409 {object} model.brandModel "Conflict"
// @response 500 {object} model.brandModel "Internal Server Error"
// @Router /api/v1/brand [get]

func (h *BrandsHandler) GetListBrands(ctx context.Context, c *gin.Context) {
	tr := otel.Tracer("component-handler")
	ctx, span := tr.Start(ctx, "BrandHandler")

	res, err := repository.GetListBrands(ctx)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	span.SetAttributes(
		attribute.Int("Count", len(res)),
	)

	c.JSON(http.StatusOK, res)

	defer span.End()
	return
}
