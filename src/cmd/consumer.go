package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

const (
	RMQ_EXCHANGETYPE    = "direct"
	RMQ_MODATATYPE      = "application/json"
	RMQ_MOEXCHANGE      = "E_MO"
	RMQ_MOQUEUE         = "Q_MO"
	RMQ_DREXCHANGE      = "E_DR"
	RMQ_DRQUEUE         = "Q_DR"
	RMQ_RENEWALEXCHANGE = "E_RENEWAL"
	RMQ_RENEWALQUEUE    = "Q_RENEWAL"
	RMQ_RETRYEXCHANGE   = "E_RETRY"
	RMQ_RETRYQUEUE      = "Q_RETRY"
	RMQ_PURGEEXCHANGE   = "E_PURGE"
	RMQ_PURGEQUEUE      = "Q_PURGE"
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
			RMQ_EXCHANGETYPE,
			true,
			RMQ_MOEXCHANGE,
			true,
			RMQ_MOQUEUE,
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			RMQ_MOQUEUE,
			RMQ_MOEXCHANGE,
			RMQ_MOQUEUE,
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
			RMQ_EXCHANGETYPE,
			true,
			RMQ_DREXCHANGE,
			true,
			RMQ_DRQUEUE,
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			RMQ_DRQUEUE,
			RMQ_DREXCHANGE,
			RMQ_DRQUEUE,
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
			RMQ_EXCHANGETYPE,
			true,
			RMQ_RENEWALEXCHANGE,
			true,
			RMQ_RENEWALQUEUE,
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			RMQ_RENEWALQUEUE,
			RMQ_RENEWALEXCHANGE,
			RMQ_RENEWALQUEUE,
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

				// 0.5 second
				time.Sleep(500 * time.Millisecond)

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
			RMQ_EXCHANGETYPE,
			true,
			RMQ_RETRYEXCHANGE,
			true,
			RMQ_RETRYQUEUE,
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			RMQ_RETRYQUEUE,
			RMQ_RETRYEXCHANGE,
			RMQ_RETRYQUEUE,
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

				// 0.5 second
				time.Sleep(500 * time.Millisecond)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerPurgeCmd = &cobra.Command{
	Use:   "consumer-purge",
	Short: "Consumer Purge Service CLI",
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
			RMQ_EXCHANGETYPE,
			true,
			RMQ_PURGEEXCHANGE,
			true,
			RMQ_PURGEQUEUE,
		)

		messagesData := queue.Rabbit.Subscribe(
			1,
			false,
			RMQ_PURGEQUEUE,
			RMQ_PURGEEXCHANGE,
			RMQ_PURGEQUEUE,
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
				purgeProccesor(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

				time.Sleep(1 * time.Second)
			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}
