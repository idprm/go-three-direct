package cmd

import (
	"encoding/json"
	"time"

	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/database"
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
			"direct",
			true,
			"E_RENEWAL",
			true,
			"Q_RENEWAL",
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
			"direct",
			true,
			"E_RETRY",
			true,
			"Q_RETRY",
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
	Use:   "publisher-purge",
	Short: "Purge CLI",
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
			"direct",
			true,
			"E_PURGE",
			true,
			"Q_PURGE",
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
		sub.PurgeAt = s.PurgeAt
		sub.IpAddress = s.IpAddress

		json, _ := json.Marshal(sub)

		queue.Rabbit.IntegratePublish(
			"E_RENEWAL",
			"Q_RENEWAL",
			"application/json",
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
		sub.PurgeAt = s.PurgeAt
		sub.IpAddress = s.IpAddress

		json, _ := json.Marshal(sub)

		queue.Rabbit.IntegratePublish(
			"E_RETRY",
			"Q_RETRY",
			"application/json",
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
		sub.PurgeAt = s.PurgeAt
		sub.IpAddress = s.IpAddress

		json, _ := json.Marshal(sub)

		queue.Rabbit.IntegratePublish(
			"E_PURGE",
			"Q_PURGE",
			"application/json",
			"",
			string(json),
		)
	}
}
