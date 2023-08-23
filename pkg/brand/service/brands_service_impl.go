package service

import (
	"context"
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/data/response"
	"golang-crud-gin/pkg/brand/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type BrandServiceImpl struct {
	BrandRepository repository.BrandRepository
	Validate        *validator.Validate
}

func NewBrandServiceImpl(brandRepository repository.BrandRepository, validate *validator.Validate) BrandService {
	return &BrandServiceImpl{
		BrandRepository: brandRepository,
		Validate:        validate,
	}
}

// FindAll implements BrandService
func (t *BrandServiceImpl) FindAll(ctx context.Context) []response.BrandResponse {
	tr := otel.Tracer("component-service")
	_, span := tr.Start(ctx, "BrandService FindAll")
	result := t.BrandRepository.FindAll(ctx)

	var res []response.BrandResponse
	for _, value := range result {
		brand := response.BrandResponse{
			Brandid: value.Brandid,
			ThBrand: value.ThBrand,
			EnBrand: value.EnBrand,
		}
		res = append(res, brand)
	}
	span.SetAttributes(
		attribute.String("StartTime", time.Now().String()),
	)
	defer span.End()
	return res
}

// FindById implements BrandService
func (t *BrandServiceImpl) FindById(brandId int, ctx context.Context) response.BrandResponse {
	tr := otel.Tracer("component-service")
	_, span := tr.Start(ctx, "brandService FindAll")
	value, err := t.BrandRepository.FindById(brandId, ctx)
	helper.ErrorPanic(err)

	brandResponse := response.BrandResponse{
		Brandid: value.Brandid,
		ThBrand: value.ThBrand,
		EnBrand: value.EnBrand,
	}
	span.SetAttributes(
		attribute.String("StartTime", time.Now().String()),
	)
	defer span.End()
	return brandResponse
}
