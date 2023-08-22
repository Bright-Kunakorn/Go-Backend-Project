package repository

import "golang-crud-gin/pkg/brand/model"

type BrandRepository interface {
	FindById(brandId int) (brand model.Brand, err error)
	FindAll() []model.Brand
}
