package route

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
	"waki.mobi/go-yatta-h3i/src/controller"
)

func Setup(app *fiber.App, db *sql.DB, gdb *gorm.DB) {

	// Default config
	app.Use(cors.New())

	// Config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	handlerMO := controller.NewHandlerMO(db, gdb)
	handlerDR := controller.NewHandlerDR(db, gdb)
	handlerIncoming := controller.NewIncomingHandler(db, gdb)

	app.Get("/moh3i", handlerMO.MessageOriginated)
	app.Get("/camph3i", handlerMO.MessageOriginated)
	app.Get("/drh3i", handlerDR.DeliveryReport)

	/**
	 * Reports
	 */
	report := app.Group("report")
	report.Get("mo", handlerIncoming.ReportMO)
	report.Get("renewal", handlerIncoming.ReportRenewal)
	report.Get("firstpush", handlerIncoming.ReportFirstpush)
}
