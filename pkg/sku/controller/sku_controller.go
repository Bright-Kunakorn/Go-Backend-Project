package controller

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/sku/data/response"
	"golang-crud-gin/pkg/sku/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type SkuController struct {
	skuService service.SkuService
}

func NewSkuController(service service.SkuService) *SkuController {
	return &SkuController{
		skuService: service,
	}
}

// FindByIdSku			godoc
// @Summary				Get Single sku by id.
// @Param				skuId path string true "update sku by id"
// @Description			Return the tahs whoes skuId valu mathes id.
// @Produce				application/json
// @Sku					sku
// @Success				200 {object} response.Response{}
// @Router				/sku/{skuId} [get]
func (controller *SkuController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid brand")
	skuId := ctx.Param("skuId")
	id, err := strconv.Atoi(skuId)
	helper.ErrorPanic(err)

	skuResponse := controller.skuService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   skuResponse,
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
func (controller *SkuController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll sku")
	skuResponse := controller.skuService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   skuResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
