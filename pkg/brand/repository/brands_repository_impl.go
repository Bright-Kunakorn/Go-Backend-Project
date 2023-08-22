package repository

import (
	"errors"

	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/model"

	"gorm.io/gorm"
)

type BrandRepositoryImpl struct {
	Db *gorm.DB
}

func NewBrandREpositoryImpl(Db *gorm.DB) BrandRepository {
	return &BrandRepositoryImpl{Db: Db}
}

// FindAll implements BrandsRepository
func (t *BrandsRepositoryImpl) FindAll() []model.Brands {
	var brand []model.Brands
	result := t.Db.Find(&brand)
	helper.ErrorPanic(result.Error)
	return brand
}

// FindById implements BrandRepository
func (t *BrandRepositoryImpl) FindById(brandId int) (brand model.Brand, err error) {
	var res model.Brand
	result := t.Db.Find(&res, brandId)
	if result != nil {
		return res, nil
	} else {
		return res, errors.New("brand is not found")
	}
}
