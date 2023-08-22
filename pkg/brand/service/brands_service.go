package service

import "golang-crud-gin/pkg/brand/data/response"

type BrandService interface {
	FindById(brandId int) response.BrandResponse
	FindAll() []response.BrandResponse
}
