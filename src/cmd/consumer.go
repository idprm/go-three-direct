package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/queue"
)

var consumerRetryCmd = &cobra.Command{
	Use:   "consumer-retry",
	Short: "Consumer Retry Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * SETUP RMQ
		 */
		queue.SetupQueue()

		/**
		 * QUEUE HANDLER
		 */
		queue.Rabbit.SetUpChannel(
			config.ViperEnv("RMQ_EXCHANGETYPE"),
			true,
			config.ViperEnv("RMQ_RETRYEXCHANGE"),
			true,
			config.ViperEnv("RMQ_RETRYQUEUE"),
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			config.ViperEnv("RMQ_RETRYQUEUE"),
			config.ViperEnv("RMQ_RETRYEXCHANGE"),
			config.ViperEnv("RMQ_RETRYQUEUE"),
		)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				// retryProccesor(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

// msisdn := c.Query("mobile_no")
// message := c.Query("message")
// code := c.Query("short_code")

// // CHECK IN BLACKLIST
// var blacklist models.Blacklist
// database.Database.Db.Where("msisdn", msisdn).First(&blacklist)

// // CHECK IN PERIODE
// var config models.Config
// database.Database.Db.Where("name", message).First(&config)

// // CHECK IN SERVICE
// var service models.Service
// database.Database.Db.Where("code", code).First(&service)

// // FILTER MESSAGE

// // SPLIT MESSAGE BY KEYWORD 1

// // SPLIT MESSAGE BY KEYWORD 2

// // SPLIT MESSAGE BY KEYWORD 3

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Worker CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
