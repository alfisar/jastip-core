package http

import "github.com/gofiber/fiber/v2"

type ProductsTravelControllerContract interface {
	Create(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
