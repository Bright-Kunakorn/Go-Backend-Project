package model

type Brand struct {
	Brandid int32   `json:"brandid"`
	ThBrand *string `json:"th_brand"`
	EnBrand *string `json:"en_brand"`
}
