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

type BrandController struct {
	brandService service.BrandService
}

func NewBrandController(service service.BrandService) *BrandController {
	return &BrandController{
		brandService: service,
	}
}

// FindByIdBrand 		godoc
// @Summary				Get Single brand by id.
// @Param				brandId path string true "update brand by id"
// @Description			Return the tahs whoes brandId valu mathes id.
// @Produce				application/json
// @Brand				brand
// @Success				200 {object} response.Response{}
// @Router				/brand/{brandId} [get]
func (controller *BrandController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid brand")
	brandId := ctx.Param("brandId")
	id, err := strconv.Atoi(brandId)
	helper.ErrorPanic(err)

	brandResponse := controller.brandService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   brandResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllBrand 		godoc
// @Summary			Get All brand.
// @Description		Return list of brand.
// @Brand			brand
// @Success			200 {obejct} response.Response{}
// @Router			/brand [get]
func (controller *BrandController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll brand")
	brandResponse := controller.brandService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   brandResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
