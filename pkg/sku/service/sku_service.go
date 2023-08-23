package service

import "golang-crud-gin/pkg/sku/data/response"

type SkuService interface {
	FindById(skuId string) response.SkuResponse
	FindAll() []response.SkuResponse
}
