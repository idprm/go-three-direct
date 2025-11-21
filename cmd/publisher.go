package cmd

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/repository"
	"github.com/idprm/go-three-direct/internal/services"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
)

var publisherRenewalCmd = &cobra.Command{
	Use:   "pub_renewal",
	Short: "Renewal CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		/**
		 * SETUP MYSQL
		 */
		db, err := connectSQL()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq := connectRabbitMq()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_RENEWAL_EXCHANGE, true, RMQ_RENEWAL_QUEUE)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {

			go func() {
				populateRenewal(db, rmq)
			}()

			time.Sleep(timeDuration * time.Minute)
		}

	},
}

var publisherRetryCmd = &cobra.Command{
	Use:   "pub_retry",
	Short: "Publisher Retry CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// connect sqldb
		db, err := connectSQL()
		if err != nil {
			panic(err)
		}

		/**
		 * connect rabbitmq
		 */
		rmq := connectRabbitMq()

		/**
		 * SETUP CHANNEL
		 */
		rmq.SetUpChannel(RMQ_EXCHANGE_TYPE, true, RMQ_RETRY_EXCHANGE, true, RMQ_RETRY_QUEUE)

		/**
		 * Looping schedule per 30 minute
		 */
		timeDuration := time.Duration(30)

		for {

			go func() {
				populateRetry(db, rmq)
			}()

			time.Sleep(timeDuration * time.Minute)
		}

	},
}

func populateRenewal(db *sql.DB, rmq rmqp.AMQP) {

	populateRepo := repository.NewPopulateRepository(db)
	populateService := services.NewPopulateService(populateRepo)

	subs := populateService.Renewal()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.Msisdn = s.Msisdn
		sub.ServiceID = s.ServiceID
		sub.Keyword = s.Keyword
		sub.PurgeAt = s.PurgeAt
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		rmq.IntegratePublish(RMQ_RENEWAL_EXCHANGE, RMQ_RENEWAL_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}

func populateRetry(db *sql.DB, rmq rmqp.AMQP) {

	populateRepo := repository.NewPopulateRepository(db)
	populateService := services.NewPopulateService(populateRepo)

	subs := populateService.Retry()

	for _, s := range *subs {
		var sub entity.Subscription

		sub.ID = s.ID
		sub.Msisdn = s.Msisdn
		sub.ServiceID = s.ServiceID
		sub.Keyword = s.Keyword
		sub.PurgeAt = s.PurgeAt
		sub.IpAddress = s.IpAddress
		sub.CreatedAt = s.CreatedAt

		json, _ := json.Marshal(sub)

		rmq.IntegratePublish(RMQ_RETRY_EXCHANGE, RMQ_RETRY_QUEUE, RMQ_DATA_TYPE, "", string(json))

		time.Sleep(100 * time.Microsecond)
	}
}
