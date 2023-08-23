package service

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/sku/data/response"
	"golang-crud-gin/pkg/sku/repository"

	"github.com/go-playground/validator/v10"
)

type SkuServiceImpl struct {
	SkuRepository repository.SkuRepository
	Validate      *validator.Validate
}

func NewSkuServiceImpl(skuRepository repository.SkuRepository, validate *validator.Validate) SkuService {
	return &SkuServiceImpl{
		SkuRepository: skuRepository,
		Validate:      validate,
	}
}

// FindAll implements BrandService
func (t *SkuServiceImpl) FindAll() []response.SkuResponse {
	result := t.SkuRepository.FindAll()

	var res []response.SkuResponse
	for _, value := range result {
		sku := response.SkuResponse{
			Skuid:           value.Skuid,
			Barcodepos:      value.Barcodepos,
			Productname:     value.Productname,
			Brandid:         value.Brandid,
			Productgroupid:  value.Productgroupid,
			Productcatid:    value.Productcatid,
			Productsubcatid: value.Productsubcatid,
			Productsizeid:   value.Productsizeid,
			Productunit:     value.Productunit,
			Packsize:        value.Packsize,
			Unit:            value.Unit,
			Banforpracharat: value.Banforpracharat,
			Isvat:           value.Isvat,
			Createby:        value.Createby,
			Createdate:      value.Createdate,
			Isactive:        value.Isactive,
			Merchantid:      value.Merchantid,
			Mapsku:          value.Mapsku,
			Isfixprice:      value.Isfixprice,
		}
		res = append(res, sku)
	}

	return res
}

// FindById implements BrandService
func (t *SkuServiceImpl) FindById(skuid string) response.SkuResponse {
	value, err := t.SkuRepository.FindById(skuid)
	helper.ErrorPanic(err)

	brandResponse := response.SkuResponse{
		Skuid:           value.Skuid,
		Barcodepos:      value.Barcodepos,
		Productname:     value.Productname,
		Brandid:         value.Brandid,
		Productgroupid:  value.Productgroupid,
		Productcatid:    value.Productcatid,
		Productsubcatid: value.Productsubcatid,
		Productsizeid:   value.Productsizeid,
		Productunit:     value.Productunit,
		Packsize:        value.Packsize,
		Unit:            value.Unit,
		Banforpracharat: value.Banforpracharat,
		Isvat:           value.Isvat,
		Createby:        value.Createby,
		Createdate:      value.Createdate,
		Isactive:        value.Isactive,
		Merchantid:      value.Merchantid,
		Mapsku:          value.Mapsku,
		Isfixprice:      value.Isfixprice,
	}
	return brandResponse
}
