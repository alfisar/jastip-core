package http

import (
	"jastip-core/application/products_travel/service"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/response"
	"github.com/gofiber/fiber/v2"
)

type productsTravelController struct {
	serv service.ProductsTravelServiceContract
}

func NewProductsTravelController(serv service.ProductsTravelServiceContract) *productsTravelController {
	return &productsTravelController{
		serv: serv,
	}
}
func (c *productsTravelController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
	return poolData
}

func (c *productsTravelController) Create(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	req := ctx.Locals("validatedData").(domain.ProductsTravelRequest)

	userID := ctx.Locals("data").(float64)

	err := c.serv.Create(ctx.Context(), poolData, int(userID), req)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessCreatedData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *productsTravelController) Delete(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	req := ctx.Locals("validatedData").(domain.ProductsTravelRequest)

	err := c.serv.Delete(ctx.Context(), poolData, req)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessCreatedData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
