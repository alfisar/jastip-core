package http

import (
	"jastip-core/application/travel_schedule/service"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/response"
	"github.com/gofiber/fiber/v2"
)

type travelController struct {
	serv service.TraveklSchServiceContract
}

func NewTravelController(serv service.TraveklSchServiceContract) *travelController {
	return &travelController{
		serv: serv,
	}
}

func (c *travelController) InitPoolData() *domain.Config {
	poolData := domain.DataPool.Get().(*domain.Config)
	return poolData
}

func (c *travelController) AddSchedule(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	req := ctx.Locals("validatedData").(domain.TravelSchRequest)

	req.UserID = int(userID)
	id, err := c.serv.AddSchedule(ctx.Context(), poolData, req)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resps := struct {
		ID int `json:"id"`
	}{
		ID: id,
	}
	resp := response.ResponseSuccess(resps, consts.SuccessCreatedData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *travelController) GetListchedule(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	params := ctx.Locals("validatedData").(domain.Params)

	totalPage, currentPage, total, result, err := c.serv.GetList(ctx.Context(), poolData, params)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccessWithPaging(result, consts.SuccessGetData, totalPage, int64(currentPage), total)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *travelController) Details(ctx *fiber.Ctx) error {
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

func (c *travelController) Update(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	id := ctx.Locals("path").(int)
	updates := ctx.Locals("validatedData").(map[string]any)

	err := c.serv.Update(ctx.Context(), poolData, int(id), int(userID), updates)
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessUpdateData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}

func (c *travelController) Delete(ctx *fiber.Ctx) error {
	poolData := c.InitPoolData()
	userID := ctx.Locals("data").(float64)
	id := ctx.Locals("path").(int)

	err := c.serv.Delete(ctx.Context(), poolData, int(id), int(userID))
	if err.Code != 0 {
		response.WriteResponse(ctx, response.Response{}, err, err.HTTPCode)
		return nil
	}

	resp := response.ResponseSuccess(nil, consts.SuccessGetData)
	response.WriteResponse(ctx, resp, err, err.Code)
	return nil
}
