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
func (t *SkuRepositoryImpl) FindAll() []model.Sku {
	var sku []model.Sku
	result := t.Db.Find(&sku)
	helper.ErrorPanic(result.Error)
	return sku
}

// FindById implements SkuRepository.
func (t *SkuRepositoryImpl) FindById(skuId string) (Sku model.Sku, err error) {
	var sku model.Sku
	result := t.Db.Where("skuid = ?", skuId).First(&sku)
	if result.Error != nil {
		return sku, result.Error
	} else if result.RowsAffected == 0 {
		return sku, errors.New("sku is not found")
	}
	return sku, nil
}

func NewSkuREpositoryImpl(Db *gorm.DB) SkuRepository {
	return &SkuRepositoryImpl{Db: Db}
}
