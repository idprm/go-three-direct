package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html"
	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/datasource/mysql/db"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
	"waki.mobi/go-yatta-h3i/src/pkg/route"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		/**
		 * LOAD CONFIG
		 */
		cfg, err := config.LoadSecret("secret.yaml")
		if err != nil {
			panic(err)
		}

		/**
		 * SETUP MYSQL
		 */
		sdb := db.InitDB(cfg)
		gdb := db.InitGormDB(cfg)

		engine := html.New("./src/presenter/views", ".html")

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
		route.Setup(cfg, app, sdb, gdb)

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
			RMQ_EXCHANGETYPE,
			true,
			RMQ_MOEXCHANGE,
			true,
			RMQ_MOQUEUE,
		)

		queue.Rabbit.SetUpChannel(
			RMQ_EXCHANGETYPE,
			true,
			RMQ_DREXCHANGE,
			true,
			RMQ_DRQUEUE,
		)

		log.Fatal(app.Listen(":" + cfg.App.Port))

	},
}
