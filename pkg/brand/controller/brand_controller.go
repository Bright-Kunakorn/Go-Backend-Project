package controller

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/data/response"
	"golang-crud-gin/pkg/brand/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type BrandsController struct {
	brandsService service.BrandsService
}

func NewBrandsController(service service.BrandsService) *BrandsController {
	return &BrandsController{
		brandsService: service,
	}
}

// FindByIdBrands 		godoc
// @Summary				Get Single brands by id.
// @Param				brandId path string true "update brands by id"
// @Description			Return the tahs whoes brandId valu mathes id.
// @Produce				application/json
// @Brands				brands
// @Success				200 {object} response.Response{}
// @Router				/brands/{brandId} [get]
func (controller *BrandsController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid brands")
	brandId := ctx.Param("brandId")
	id, err := strconv.Atoi(brandId)
	helper.ErrorPanic(err)

	brandResponse := controller.brandsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   brandResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllBrands 		godoc
// @Summary			Get All brands.
// @Description		Return list of brands.
// @Brands			brands
// @Success			200 {obejct} response.Response{}
// @Router			/brands [get]
func (controller *BrandsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll brands")
	brandResponse := controller.brandsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   brandResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
