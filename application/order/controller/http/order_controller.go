package http

import (
	"jastip-core/application/order/service"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/response"
	"github.com/gofiber/fiber/v2"
)

type orderController struct {
	serv service.OrderServiceContract
}

func NewOrderService(serv service.OrderServiceContract) *orderController {
	return &orderController{
		serv: serv,
	}
}

func (c *orderController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
	return poolData
}

func (c *orderController) Create(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	req := ctx.Locals("validatedData").(domain.OrderRequest)

	err := c.serv.Create(ctx.Context(), poolData, req, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessCreatedData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *orderController) GetList(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	params := ctx.Locals("validatedData").(domain.Params)

	totalPage, currentPage, total, result, err := c.serv.GetList(ctx.Context(), poolData, params, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccessWithPaging(result, consts.SuccessGetData, totalPage, int64(currentPage), total)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *orderController) Details(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	id := ctx.Locals("path").(int)

	result, err := c.serv.GetDetails(ctx.Context(), poolData, int(id), int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(result, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
