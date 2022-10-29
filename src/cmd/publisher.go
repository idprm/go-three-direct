package cmd

import (
	"encoding/json"
	"time"

	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
	"waki.mobi/go-yatta-h3i/src/pkg/query"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

var publisherRenewalCmd = &cobra.Command{
	Use:   "publisher-renewal",
	Short: "Renewal CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

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
			config.ViperEnv("RMQ_RENEWALEXCHANGE"),
			true,
			config.ViperEnv("RMQ_RENEWALQUEUE"),
		)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			currentTime := time.Now()
			timeNow := currentTime.Format("15:04")

			var schedule model.Schedule
			resultPublish := database.Datasource.DB().
				Where("name", "RENEWAL_PUSH").
				Where("TIME(publish_at) = TIME(?)", timeNow).
				Where("status", true).
				First(&schedule)

			resultLocked := database.Datasource.DB().
				Where("name", "RENEWAL_PUSH").
				Where("TIME(un_locked_at) = TIME(?)", timeNow).
				Where("status", false).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				database.Datasource.DB().Save(&schedule)

				go func() {
					populateRenewal()
				}()

			}

			if resultLocked.RowsAffected == 1 {
				schedule.Status = true
				database.Datasource.DB().Save(&schedule)
			}

			time.Sleep(timeDuration * time.Minute)

		}

	},
}

var publisherRetryCmd = &cobra.Command{
	Use:   "publisher-retry",
	Short: "Retry CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

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
			config.ViperEnv("RMQ_RETRYEXCHANGE"),
			true,
			config.ViperEnv("RMQ_RETRYQUEUE"),
		)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			currentTime := time.Now()
			timeNow := currentTime.Format("15:04")

			var schedule model.Schedule
			resultPublish := database.Datasource.DB().
				Where("name", "RETRY_PUSH").
				Where("TIME(publish_at) = TIME(?)", timeNow).
				Where("status", true).
				First(&schedule)

			resultLocked := database.Datasource.DB().
				Where("name", "RETRY_PUSH").
				Where("TIME(un_locked_at) = TIME(?)", timeNow).
				Where("status", false).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				database.Datasource.DB().Save(&schedule)

				go func() {
					populateRetry()
				}()

			}

			if resultLocked.RowsAffected == 1 {
				schedule.Status = true
				database.Datasource.DB().Save(&schedule)
			}

			time.Sleep(timeDuration * time.Minute)

		}

	},
}

var publisherPurgeCmd = &cobra.Command{
	Use:   "publisher-retry",
	Short: "Retry CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

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
			config.ViperEnv("RMQ_RETRYEXCHANGE"),
			true,
			config.ViperEnv("RMQ_RETRYQUEUE"),
		)

		/**
		 * Looping schedule
		 */
		timeDuration := time.Duration(1)

		for {
			currentTime := time.Now()
			timeNow := currentTime.Format("15:04")

			var schedule model.Schedule
			resultPublish := database.Datasource.DB().
				Where("name", "PURGE_PUSH").
				Where("TIME(publish_at) = TIME(?)", timeNow).
				Where("status", true).
				First(&schedule)

			resultLocked := database.Datasource.DB().
				Where("name", "PURGE_PUSH").
				Where("TIME(un_locked_at) = TIME(?)", timeNow).
				Where("status", false).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				database.Datasource.DB().Save(&schedule)

				go func() {
					populatePurge()
				}()

			}

			if resultLocked.RowsAffected == 1 {
				schedule.Status = true
				database.Datasource.DB().Save(&schedule)
			}

			time.Sleep(timeDuration * time.Minute)

		}

	},
}

func populateRenewal() {

	subs, _ := query.GetDataPopulate("RENEWAL")

	for _, s := range subs {
		var sub model.Subscription

		sub.ID = s.ID
		sub.Msisdn = s.Msisdn
		sub.ServiceID = s.ServiceID
		sub.Keyword = s.Keyword
		sub.IpAddress = s.IpAddress

		json, _ := json.Marshal(sub)

		queue.Rabbit.IntegratePublish(
			config.ViperEnv("RMQ_RENEWALEXCHANGE"),
			config.ViperEnv("RMQ_RENEWALQUEUE"),
			config.ViperEnv("RMQ_RENEWALDATATYPE"),
			"",
			string(json),
		)
	}

}

func populateRetry() {

	subs, _ := query.GetDataPopulate("RETRY")

	for _, s := range subs {
		var sub model.Subscription

		sub.ID = s.ID
		sub.Msisdn = s.Msisdn
		sub.ServiceID = s.ServiceID
		sub.Keyword = s.Keyword
		sub.IpAddress = s.IpAddress

		json, _ := json.Marshal(sub)

		queue.Rabbit.IntegratePublish(
			config.ViperEnv("RMQ_RETRYEXCHANGE"),
			config.ViperEnv("RMQ_RETRYQUEUE"),
			config.ViperEnv("RMQ_RETRYDATATYPE"),
			"",
			string(json),
		)
	}
}

func populatePurge() {
	subs, _ := query.GetDataPopulate("PURGE")

	for _, s := range subs {
		var sub model.Subscription

		sub.ID = s.ID
		sub.Msisdn = s.Msisdn
		sub.ServiceID = s.ServiceID
		sub.Keyword = s.Keyword
		sub.IpAddress = s.IpAddress

		json, _ := json.Marshal(sub)

		queue.Rabbit.IntegratePublish(
			config.ViperEnv("RMQ_PURGEEXCHANGE"),
			config.ViperEnv("RMQ_PURGEQUEUE"),
			config.ViperEnv("RMQ_PURGEDATATYPE"),
			"",
			string(json),
		)
	}
}
