package router

import "github.com/gofiber/fiber/v2"

func NewRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:   false,
		BodyLimit: 20 * 1024 * 1024,
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		return c.Next()
	})

	api := app.Group("api/v1")
	api.Options("/*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		return c.SendStatus(fiber.StatusOK)
	})

	SimpleInit().simpleRouters(api)
	TravelSchInit().travelSchRouters(api)
	CountriesInit().countriesRouters(api)
	ProductsInit().productRouters(api)
	ProductsTravelInit().productTravelRouters(api)
	return app
}
