package controller

import (
	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/helper"
	"github.com/alfisar/jastip-import/helpers/response"
	"github.com/gofiber/fiber/v2"
)

type simpleController struct {
}

func NewSimpleController() *simpleController {
	return &simpleController{}
}

func (c *simpleController) Healthy(ctx *fiber.Ctx) error {

	response.WriteResponse(ctx, response.Response{
		Status:  "Success",
		Code:    0,
		Message: "Welcome to core jastip api :)",
		MetaData: response.MetaData{
			Timestamp: helper.TimeGenerator(),
			Version:   "v1",
		},
	}, domain.ErrorData{}, 200)
	return nil
}
