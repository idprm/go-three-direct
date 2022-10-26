package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

var consumerMOCmd = &cobra.Command{
	Use:   "consumer-mo",
	Short: "Consumer MO Service CLI",
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
			config.ViperEnv("RMQ_MOEXCHANGE"),
			true,
			config.ViperEnv("RMQ_MOQUEUE"),
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			config.ViperEnv("RMQ_MOQUEUE"),
			config.ViperEnv("RMQ_MOEXCHANGE"),
			config.ViperEnv("RMQ_MOQUEUE"),
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
				moProccesor(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

				time.Sleep(1 * time.Millisecond)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerDRCmd = &cobra.Command{
	Use:   "consumer-dr",
	Short: "Consumer DR Service CLI",
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
			config.ViperEnv("RMQ_DREXCHANGE"),
			true,
			config.ViperEnv("RMQ_DRQUEUE"),
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			config.ViperEnv("RMQ_DRQUEUE"),
			config.ViperEnv("RMQ_DREXCHANGE"),
			config.ViperEnv("RMQ_DRQUEUE"),
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
				drProccesor(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

				time.Sleep(1 * time.Millisecond)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerRenewalCmd = &cobra.Command{
	Use:   "consumer-renewal",
	Short: "Consumer Renewal Service CLI",
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
			config.ViperEnv("RMQ_RENEWALEXCHANGE"),
			true,
			config.ViperEnv("RMQ_RENEWALQUEUE"),
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			config.ViperEnv("RMQ_RENEWALQUEUE"),
			config.ViperEnv("RMQ_RENEWALEXCHANGE"),
			config.ViperEnv("RMQ_RENEWALQUEUE"),
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
				renewalProccesor(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

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
				retryProccesor(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}
