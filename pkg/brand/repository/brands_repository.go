package repository

import "golang-crud-gin/pkg/brand/model"

type BrandsRepository interface {
	FindById(brandsId int) (brands model.Brands, err error)
	FindAll() []model.Brands
}
