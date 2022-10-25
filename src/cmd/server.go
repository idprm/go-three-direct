package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/route"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Setup routing rules
		route := route.SetupRouter()

		// Setup Trusted IP
		route.SetTrustedProxies([]string{"192.168.1.2"})

		// Logger
		route.Use(gin.Logger())

		/**
		 * Access log on browser
		 */
		route.StaticFS("/logs", http.Dir("logs"))

		route.Run(":" + config.ViperEnv("APP_PORT"))
	},
}
