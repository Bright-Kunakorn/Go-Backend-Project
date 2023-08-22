package repository

import (
	"errors"

	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/sku/model"

	"gorm.io/gorm"
)

type SkuRepositoryImpl struct {
	Db *gorm.DB
}

// FindAll implements SkuRepository.
func (t *SkuRepositoryImpl) FindAll() []model.Backendposdatasku {
	var sku []model.Backendposdatasku
	result := t.Db.Find(&sku)
	helper.ErrorPanic(result.Error)
	return sku
}

// FindById implements SkuRepository.
func (t *SkuRepositoryImpl) FindById(skuId int) (backendposdatasku model.Backendposdatasku, err error) {
	var sku model.Backendposdatasku
	result := t.Db.Find(&sku, skuId)
	if result != nil {
		return sku, nil
	} else {
		return sku, errors.New("sku is not found")
	}
}

func NewSkuREpositoryImpl(Db *gorm.DB) SkuRepository {
	return &SkuRepositoryImpl{Db: Db}
}
