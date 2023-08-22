package service

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/data/response"
	"golang-crud-gin/pkg/brand/repository"

	"github.com/go-playground/validator/v10"
)

type BrandsServiceImpl struct {
	BrandsRepository repository.BrandsRepository
	Validate         *validator.Validate
}

func NewBrandsServiceImpl(brandRepository repository.BrandsRepository, validate *validator.Validate) BrandsService {
	return &BrandsServiceImpl{
		BrandsRepository: brandRepository,
		Validate:         validate,
	}
}

// FindAll implements BrandsService
func (t *BrandsServiceImpl) FindAll() []response.BrandsResponse {
	result := t.BrandsRepository.FindAll()

	var brands []response.BrandsResponse
	for _, value := range result {
		brand := response.BrandsResponse{
			Brandid: value.Brandid,
			ThBrand: value.ThBrand,
			EnBrand: value.EnBrand,
		}
		brands = append(brands, brand)
	}

	return brands
}

// FindById implements BrandsService
func (t *BrandsServiceImpl) FindById(brandsId int) response.BrandsResponse {
	value, err := t.BrandsRepository.FindById(brandsId)
	helper.ErrorPanic(err)

	brandResponse := response.BrandsResponse{
		Brandid: value.Brandid,
		ThBrand: value.ThBrand,
		EnBrand: value.EnBrand,
	}
	return brandResponse
}
