package repository

import (
	"context"
	"golang-crud-gin/pkg/brand/model"
)

type BrandRepository interface {
	FindById(brandId int,ctx context.Context) (brand model.Brand, err error)
	FindAll(ctx context.Context) []model.Brand
}
