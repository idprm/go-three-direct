package app

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/wiliehidayat87/rmqp"
	"waki.mobi/go-yatta-h3i/src/config"
)

func mapUrls(cfg *config.Secret, db *sql.DB, rmpq rmqp.AMQP) *fiber.App {
	engine := html.New("./src/presenter/views", ".html")

	/**
	 * Init Fiber
	 */
	router := fiber.New(fiber.Config{
		Views: engine,
	})

	/**
	 * Access log on browser
	 */
	router.Use("/logs", filesystem.New(filesystem.Config{
		Root:         http.Dir(cfg.Log.Path),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "404.html",
		MaxAge:       3600,
	}))

	/**
	 * Write access logger
	 */
	// file, err := os.OpenFile(cfg.Log.Path+"/access_log/log-"+time.Now().Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }

	// router.Use(requestid.New())
	// router.Use(access_logger.New(access_logger.Config{
	// 	Format:     "${pid} ${status} - ${method} ${path}\n",
	// 	TimeFormat: "02-Jan-2006",
	// 	TimeZone:   cfg.App.TimeZone,
	// 	Output:     file,
	// }))

	// path, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }

	// serviceRepo := repository.NewServiceRepository(db)
	// serviceService := services.NewServiceService(serviceRepo)

	// transactionRepo := repository.NewTransactionRepository(db)
	// transactionService := services.NewTransactionService(transactionRepo)

	return router
}
