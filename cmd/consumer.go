package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/spf13/cobra"
)

const (
	RMQ_EXCHANGE_TYPE    string = "direct"
	RMQ_DATA_TYPE        string = "application/json"
	RMQ_MO_EXCHANGE      string = "E_MO"
	RMQ_MO_QUEUE         string = "Q_MO"
	RMQ_DR_EXCHANGE      string = "E_DR"
	RMQ_DR_QUEUE         string = "Q_DR"
	RMQ_RENEWAL_EXCHANGE string = "E_RENEWAL"
	RMQ_RENEWAL_QUEUE    string = "Q_RENEWAL"
	RMQ_RETRY_EXCHANGE   string = "E_RETRY"
	RMQ_RETRY_QUEUE      string = "Q_RETRY"
	RMQ_NOTIF_EXCHANGE   string = "E_NOTIF"
	RMQ_NOTIF_QUEUE      string = "Q_NOTIF"
	RMQ_PURGE_EXCHANGE   string = "E_PURGE"
	RMQ_PURGE_QUEUE      string = "Q_PURGE"
)

var consumerMOCmd = &cobra.Command{
	Use:   "mo",
	Short: "Consumer MO Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDB()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq := connectRabbitMq()

		l := logger.NewLogger()

		/**
		 * QUEUE HANDLER
		 */
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_MO_EXCHANGE,
			true,
			RMQ_MO_QUEUE,
		)

		messagesData := rmq.Subscribe(
			1,
			false,
			RMQ_MO_QUEUE,
			RMQ_MO_EXCHANGE,
			RMQ_MO_QUEUE,
		)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, l)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.MO(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

				// 1 second
				time.Sleep(1 * time.Second)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerDRCmd = &cobra.Command{
	Use:   "dr",
	Short: "Consumer Delivery Notification Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		db, err := connectDB()
		if err != nil {
			panic(err)
		}

		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		rmq := connectRabbitMq()

		l := logger.NewLogger()

		/**
		 * QUEUE HANDLER
		 */
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_DR_EXCHANGE,
			true,
			RMQ_DR_QUEUE,
		)

		messagesData := rmq.Subscribe(
			1,
			false,
			RMQ_DR_QUEUE,
			RMQ_DR_EXCHANGE,
			RMQ_DR_QUEUE,
		)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, l)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.DR(&wg, d.Body)
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
	Use:   "renewal",
	Short: "Consumer Renewal Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect database
		 */
		db, err := connectDB()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq := connectRabbitMq()

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		l := logger.NewLogger()

		/**
		 * QUEUE HANDLER
		 */
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_RENEWAL_EXCHANGE,
			true,
			RMQ_RENEWAL_QUEUE,
		)

		messagesData := rmq.Subscribe(
			1,
			false,
			RMQ_RENEWAL_QUEUE,
			RMQ_RENEWAL_EXCHANGE,
			RMQ_RENEWAL_QUEUE,
		)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, l)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.Renewal(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

				// 0.22 second
				time.Sleep(220 * time.Millisecond)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerRetryCmd = &cobra.Command{
	Use:   "retry",
	Short: "Consumer Retry Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDB()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq := connectRabbitMq()

		l := logger.NewLogger()

		/**
		 * QUEUE HANDLER
		 */
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_RETRY_EXCHANGE,
			true,
			RMQ_RETRY_QUEUE,
		)

		messagesData := rmq.Subscribe(
			1,
			false,
			RMQ_RETRY_QUEUE,
			RMQ_RETRY_EXCHANGE,
			RMQ_RETRY_QUEUE,
		)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, l)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.Retry(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

				// 0.19 second
				time.Sleep(190 * time.Millisecond)
			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}

var consumerNotif = &cobra.Command{
	Use:   "notif",
	Short: "Consumer Notification Service CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * connect mysql
		 */
		db, err := connectDB()
		if err != nil {
			panic(err)
		}

		/**
		 * connect redis
		 */
		rds, err := connectRedis()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq := connectRabbitMq()

		l := logger.NewLogger()

		/**
		 * QUEUE HANDLER
		 */
		rmq.SetUpChannel(
			RMQ_EXCHANGE_TYPE,
			true,
			RMQ_NOTIF_EXCHANGE,
			true,
			RMQ_NOTIF_QUEUE,
		)

		messagesData := rmq.Subscribe(
			1,
			false,
			RMQ_NOTIF_QUEUE,
			RMQ_NOTIF_EXCHANGE,
			RMQ_NOTIF_QUEUE,
		)

		// Initial sync waiting group
		var wg sync.WaitGroup

		// Loop forever listening incoming data
		forever := make(chan bool)

		p := NewProcessor(db, rds, l)

		// Set into goroutine this listener
		go func() {

			// Loop every incoming data
			for d := range messagesData {

				wg.Add(1)
				p.Notif(&wg, d.Body)
				wg.Wait()

				// Manual consume queue
				d.Ack(false)

				// 1 second
				time.Sleep(1 * time.Second)

			}

		}()

		fmt.Println("[*] Waiting for data...")

		<-forever
	},
}
