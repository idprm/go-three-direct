package cmd

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

		/**
		 * Init Fiber
		 */
		app := fiber.New()

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
