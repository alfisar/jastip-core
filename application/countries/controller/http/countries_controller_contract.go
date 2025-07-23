package http

import "github.com/gofiber/fiber/v2"

type CountriesConrollerContract interface {
	GetList(ctx *fiber.Ctx) error
}
