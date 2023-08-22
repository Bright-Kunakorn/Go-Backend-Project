package repository

import "golang-crud-gin/pkg/sku/model"

type SkuRepository interface {
	FindById(skuId int) (backendposdatasku model.Backendposdatasku, err error)
	FindAll() []model.Backendposdatasku
}
