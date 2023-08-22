package service

import (
	"golang-crud-gin/data/response"
)

type BrandsService interface {
	FindById(brandsId int) response.BrandsResponse
	FindAll() []response.BrandsResponse
}
