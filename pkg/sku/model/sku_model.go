package model

import "time"

type Sku struct {
	Skuid           string    `json:"skuid"`
	Barcodepos      *string   `json:"barcodepos"`
	Productname     *string   `json:"productname"`
	Brandid         *int32    `json:"brandid"`
	Productgroupid  *int32    `json:"productgroupid"`
	Productcatid    *int32    `json:"productcatid"`
	Productsubcatid *int32    `json:"productsubcatid"`
	Productsizeid   *int32    `json:"productsizeid"`
	Productunit     *int32    `json:"productunit"`
	Packsize        *string   `json:"packsize"`
	Unit            *int32    `json:"unit"`
	Banforpracharat *int32    `json:"banforpracharat"`
	Isvat           *int16    `json:"isvat"`
	Createby        *string   `json:"createby"`
	Createdate      time.Time `json:"createdate"`
	Isactive        int16     `json:"isactive"`
	Merchantid      *string   `json:"merchantid"`
	Mapsku          *string   `json:"mapsku"`
	Isfixprice      int16     `json:"isfixprice"`
}
