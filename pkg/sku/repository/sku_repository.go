package repository

import "golang-crud-gin/pkg/sku/model"

type BrandsRepository interface {
	FindById(brandsId int) (brands model.sku, err error)
	FindAll() []model.Sku
}
