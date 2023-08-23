package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/model"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
)

type BrandRepositoryImpl struct {
	Db *gorm.DB
}

func NewBrandREpositoryImpl(Db *gorm.DB) BrandRepository {
	return &BrandRepositoryImpl{Db: Db}
}

// FindAll implements BrandRepository
func (t *BrandRepositoryImpl) FindAll(ctx context.Context) []model.Brand {
	tr := otel.Tracer("component-repository")
	_, span := tr.Start(ctx, "BrandRepository FindAll")
	var brand []model.Brand
	result := t.Db.Find(&brand)
	helper.ErrorPanic(result.Error)
	span.SetAttributes(
		attribute.String("SQLCommand", "SELECT * FROM brand"),
		attribute.String("SQLQuery", "findbyid brand"),
		attribute.String("SQLTime", time.Now().String()),
	)
	defer span.End()
	return brand
}

// FindById implements BrandRepository
func (t *BrandRepositoryImpl) FindById(brandId int, ctx context.Context) (brand model.Brand, err error) {
	tr := otel.Tracer("component-repository")
	_, span := tr.Start(ctx, "BrandRepository FindById")

	defer span.End()
	var res model.Brand
	result := t.Db.Find(&res, brandId)
	if result != nil {
		span.SetAttributes(
			attribute.String("SQLCommand", "SELECT * FROM brand WHERE brand_id = "+strconv.Itoa(brandId)),
			attribute.String("SQLQuery", "findbyid brand"),
			attribute.String("SQLTime", time.Now().String()),
			attribute.String("IsError", result.Error.Error()),
		)
		return res, nil
	} else {
		span.SetAttributes(
			attribute.String("SQLCommand", "SELECT * FROM brand WHERE brand_id = "+strconv.Itoa(brandId)),
			attribute.String("SQLQuery", "findbyid brand"),
			attribute.String("SQLTime", time.Now().String()),
			attribute.String("IsError", "brand is not found"),
		)
		return res, errors.New("brand is not found")
	}

}
