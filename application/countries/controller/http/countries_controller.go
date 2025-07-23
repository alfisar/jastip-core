package http

import (
	"jastip-core/application/countries/service"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/response"
	"github.com/gofiber/fiber/v2"
)

type countriesController struct {
	serv service.CountriesServiceContract
}

func NewCountriesController(serv service.CountriesServiceContract) *countriesController {
	return &countriesController{
		serv: serv,
	}
}

func (c *countriesController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
	return poolData
}

func (c *countriesController) GetList(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	params := ctx.Locals("validatedData").(domain.Params)

	result, currentPage, total, limit, totalPage, err := c.serv.GetList(poolData, params)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccessWithFullPaging(result, consts.SuccessGetData, totalPage, int64(currentPage), total, limit)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
