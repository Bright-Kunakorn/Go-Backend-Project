package repository

import "golang-crud-gin/pkg/sku/model"

type SkuRepository interface {
	FindById(skuId string) (Sku model.Sku, err error)
	FindAll() []model.Sku
}
