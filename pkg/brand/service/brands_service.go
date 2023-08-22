package service

import "golang-crud-gin/pkg/brand/data/response"

type BrandsService interface {
	FindById(brandsId int) response.BrandsResponse
	FindAll() []response.BrandsResponse
}
