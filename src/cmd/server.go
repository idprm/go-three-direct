package cmd

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

		// Setup routing rules
		router := route.SetupRouter()

		// Setup Trusted IP
		// route.SetTrustedProxies([]string{"192.168.1.2"})

		// Logger
		router.Use(gin.Logger())

		/**
		 * Access log on browser
		 */
		router.StaticFS("/logs", http.Dir("logs"))

		server := &http.Server{
			Addr:           ":" + config.ViperEnv("APP_PORT"),
			Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

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
		server.ListenAndServe()
	},
}
