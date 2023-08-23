package service

import (
	"context"
	"golang-crud-gin/pkg/brand/data/response"
)

type BrandService interface {
	FindById(brandId int,ctx context.Context) response.BrandResponse
	FindAll(ctx context.Context) []response.BrandResponse
}
