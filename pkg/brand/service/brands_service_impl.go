package service

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/data/response"
	"golang-crud-gin/pkg/brand/repository"

	"github.com/go-playground/validator/v10"
)

type BrandServiceImpl struct {
	BrandRepository repository.BrandRepository
	Validate        *validator.Validate
}

func NewBrandServiceImpl(brandRepository repository.BrandRepository, validate *validator.Validate) BrandService {
	return &BrandServiceImpl{
		BrandRepository: brandRepository,
		Validate:        validate,
	}
}

// FindAll implements BrandService
func (t *BrandServiceImpl) FindAll() []response.BrandResponse {
	result := t.BrandRepository.FindAll()

	var res []response.BrandResponse
	for _, value := range result {
		brand := response.BrandResponse{
			Brandid: value.Brandid,
			ThBrand: value.ThBrand,
			EnBrand: value.EnBrand,
		}
		res = append(res, brand)
	}

	return res
}

// FindById implements BrandService
func (t *BrandServiceImpl) FindById(brandId int) response.BrandResponse {
	value, err := t.BrandRepository.FindById(brandId)
	helper.ErrorPanic(err)

	brandResponse := response.BrandResponse{
		Brandid: value.Brandid,
		ThBrand: value.ThBrand,
		EnBrand: value.EnBrand,
	}
	return brandResponse
}
