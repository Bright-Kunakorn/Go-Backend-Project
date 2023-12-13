package controller

import (
	"golang-crud-gin/helper"
	"golang-crud-gin/pkg/brand/data/response"
	"golang-crud-gin/pkg/brand/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type BrandController struct {
	brandService service.BrandService
}

func NewBrandController(service service.BrandService) *BrandController {
	return &BrandController{
		brandService: service,
	}
}

// func (r CustomerRepository) UpdateCustomerInfo(supplierId string, customerInfo customermodel.CustomerInfoResponse) error {
//     ctx, cancel := r.withTimeout()
//     defer cancel()

//     if err := r.Database.WithContext(ctx).
//         Table(model.TableNameContract).
//         Joins("JOIN " + model.TableNameCustomer + " ON " + model.TableNameContract + ".customer_id = " + model.TableNameCustomer + ".id").
//         Where(model.TableNameContract+".customer_id = ? AND "+model.TableNameContract+".supplier_id = ?", customerInfo.CustomerId, supplierId).
//         Updates(map[string]interface{}{
//             model.TableNameContract + ".price_tier_group_id": customerInfo.PriceTierGroupId,
//             model.TableNameCustomer + ".name":              customerInfo.Name,
//             model.TableNameCustomer + ".phone_1":          customerInfo.Phone1,
//             model.TableNameCustomer + ".phone_2":          customerInfo.Phone2,
//             model.TableNameContract + ".updated_by":       customerInfo.UpdatedBy,
//             model.TableNameCustomer + ".updated_by":       customerInfo.UpdatedBy,
//             model.TableNameContract + ".updated_date":     customerInfo.UpdatedDate,
//             model.TableNameCustomer + ".updated_date":     customerInfo.UpdatedDate,
//         }).
//         Error; err != nil {
//         return err
//     }

//     return nil
// }

// FindByIdBrand 		godoc
// @Summary				Get Single brand by id.
// @Param				brandId path string true "update brand by id"
// @Description			Return the tahs whoes brandId valu mathes id.
// @Produce				application/json
// @Brand				brand
// @Success				200 {object} response.Response{}
// @Router				/brand/{brandId} [get]
func (controller *BrandController) FindById(ctx *gin.Context) {
	tr := otel.Tracer("component-controller")
	ctxWithTrace, span := tr.Start(ctx, "BrandController FindById")
	log.Info().Msg("findbyid brand")

	brandId := ctx.Param("brandId")
	id, err := strconv.Atoi(brandId)
	helper.ErrorPanic(err)

	brandResponse := controller.brandService.FindById(id, ctxWithTrace)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   brandResponse,
	}
	spanAttributes := []attribute.KeyValue{
		attribute.String("StartTime", time.Now().String()),
	}

	if err != nil {
		spanAttributes = append(spanAttributes, attribute.String("IsError", err.Error()))
	}

	span.SetAttributes(spanAttributes...)
	defer span.End()

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}


// FindAllBrand 	godoc
// @Summary			Get All brand.
// @Description		Return list of brand.
// @Brand			brand
// @Success			200 {obejct} response.Response{}
// @Router			/brand [get]
func (controller *BrandController) FindAll(ctx *gin.Context) {
	tr := otel.Tracer("component-controller")
	ctxWithTrace, span := tr.Start(ctx, "BrandController FindById")
	log.Info().Msg("findAll brand")
	brandResponse := controller.brandService.FindAll(ctxWithTrace)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   brandResponse,
	}

	span.SetAttributes(
		attribute.String("StartTime", time.Now().String()),
	)
	defer span.End()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
