package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"waki.mobi/go-yatta-h3i/src/controller"
)

func Setup(app *fiber.App) {

	// Default config
	app.Use(cors.New())

	// Config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, https://h3i-linkit.vercel.app",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/moh3i", controller.MessageOriginated)
	app.Get("/drh3i", controller.DeliveryReport)
	app.Get("/testmoh3i", controller.TestMO)
	app.Get("/testdr3i", controller.TestDR)
}
