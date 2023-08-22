package repository

import (
	"errors"
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/model"

	"gorm.io/gorm"
)

type BrandsRepositoryImpl struct {
	Db *gorm.DB
}

func NewBrandsREpositoryImpl(Db *gorm.DB) BrandsRepository {
	return &BrandsRepositoryImpl{Db: Db}
}

// FindAll implements BrandsRepository
func (t *BrandsRepositoryImpl) FindAll() []model.Brands {
	var brand []model.Brands
	result := t.Db.Find(&brand)
	helper.ErrorPanic(result.Error)
	return brand
}

// FindById implements BrandsRepository
func (t *BrandsRepositoryImpl) FindById(brandsId int) (brands model.Brands, err error) {
	var brand model.Brands
	result := t.Db.Find(&brand, brandsId)
	if result != nil {
		return brand, nil
	} else {
		return brand, errors.New("brand is not found")
	}
}
