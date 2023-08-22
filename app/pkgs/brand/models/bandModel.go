package models

type Brands struct {
	Brands []Brand `json:"brands"`
}

type Brand struct {
	Brandid int32   `json:"brandid" binding:"required" example:"1" maxLength:"255"`       // Brand ID
	ThBrand *string `json:"th_brand" binding:"required" example:"string" maxLength:"255"` // Brand Th Name
	EnBrand *string `json:"en_name" binding:"required" example:"string" maxLength:"255"`  // Brand En name
}
