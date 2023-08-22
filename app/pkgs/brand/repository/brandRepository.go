package repository

import (
	"context"

	"github.com/shall-we-go/go-gin-swagger-example/database"
	"github.com/shall-we-go/go-gin-swagger-example/pkgs/brand/models"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func GetListBrands(ctx context.Context) ([]models.Brand, error) {
	tr := otel.Tracer("component-repository")
	ctx, span := tr.Start(ctx, "BrandRepository")

	var brands []models.Brand

	res := database.DB.Table("brand").Find(&brands)
	if res.Error != nil {
		return nil, res.Error
	}

	span.SetAttributes(
		attribute.Int("Count", len(brands)),
		attribute.String("SQLCommand", res.Statement.SQL.String()),
	)
	defer span.End()
	return brands, nil
}
