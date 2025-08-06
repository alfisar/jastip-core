package http

import (
	"jastip-core/application/products/service"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/response"
	"github.com/gofiber/fiber/v2"
)

type productsController struct {
	serv service.ProductsServiceContract
}

func NewProductController(serv service.ProductsServiceContract) *productsController {
	return &productsController{
		serv: serv,
	}
}

func (c *productsController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
	return poolData
}

func (c *productsController) Create(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	req := ctx.Locals("validatedData").(domain.ProductData)

	req.UserID = int(userID)
	r := ctx.Request()
	err := c.serv.Create(ctx.Context(), poolData, req, r)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessCreatedData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *productsController) GetList(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	params := ctx.Locals("validatedData").(domain.Params)

	totalPage, currentPage, total, limit, result, err := c.serv.GetList(poolData, int(userID), params)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccessWithFullPaging(result, consts.SuccessGetData, totalPage, int64(currentPage), total, limit)
	response.WriteResponse(ctx, resp, err, fiber.StatusOK)
	return nil
}

func (c *productsController) GetListProductTravel(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	travelID := ctx.Locals("path").(int)
	params := ctx.Locals("validatedData").(domain.Params)

	totalPage, currentPage, total, limit, result, err := c.serv.GetListProductTravel(poolData, int(userID), travelID, params)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccessWithFullPaging(result, consts.SuccessGetData, totalPage, int64(currentPage), total, limit)
	response.WriteResponse(ctx, resp, err, fiber.StatusOK)
	return nil
}

func (c *productsController) Update(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	id := ctx.Locals("path").(int)
	updates := ctx.Locals("validatedData").(map[string]any)
	file, _ := ctx.MultipartForm()

	err := c.serv.Update(ctx.Context(), poolData, id, int(userID), updates, file)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, fiber.StatusOK)
	return nil
}

func (c *productsController) Delete(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	id := ctx.Locals("path").(int)

	err := c.serv.Delete(poolData, id, int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, fiber.StatusOK)
	return nil
}
