package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
	"waki.mobi/go-yatta-h3i/src/pkg/route"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		engine := html.New("./src/views", ".html")

		/**
		 * Init Fiber
		 */
		app := fiber.New(fiber.Config{
			Views: engine,
		})

		/**
		 * Access log on browser
		 */
		app.Use("/logs", filesystem.New(filesystem.Config{
			Root:         http.Dir("./logs"),
			Browse:       true,
			Index:        "index.html",
			NotFoundFile: "404.html",
			MaxAge:       3600,
		}))

		/**
		 * SETUP route
		 */
		route.Setup(app)

		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		app.Static("/static", path+"/public")

		/**
		 * SETUP RMQ
		 */
		queue.SetupQueue()

		/**
		 * SETUP CHANNEL
		 */
		queue.Rabbit.SetUpChannel(
			config.ViperEnv("RMQ_EXCHANGETYPE"),
			true,
			config.ViperEnv("RMQ_MOEXCHANGE"),
			true,
			config.ViperEnv("RMQ_MOQUEUE"),
		)

		queue.Rabbit.SetUpChannel(
			config.ViperEnv("RMQ_EXCHANGETYPE"),
			true,
			config.ViperEnv("RMQ_DREXCHANGE"),
			true,
			config.ViperEnv("RMQ_DRQUEUE"),
		)

		log.Fatal(app.Listen(":" + config.ViperEnv("APP_PORT")))

	},
}
