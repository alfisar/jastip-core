package controller

import (
	"github.com/alfisar/jastip-import/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type simpleController struct {
}

func NewSimpleController() *simpleController {
	return &simpleController{}
}

func (c *simpleController) Healthy(ctx *fiber.Ctx) error {

	_ = domain.DataPool
	ctx.Status(fasthttp.StatusOK).JSON(domain.ErrorData{
		Status:  "success",
		Code:    0,
		Message: "Welcome to API Core Justip.in version 1.0, enjoy and chersss :)",
	})
	return nil
}
