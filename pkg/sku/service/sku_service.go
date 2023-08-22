package service

import "golang-crud-gin/pkg/sku/data/response"

type SkuService interface {
	FindById(brandId int) response.SkuResponse
	FindAll() []response.SkuResponse
}
