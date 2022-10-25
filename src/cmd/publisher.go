package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
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
				Where("status", 1).
				First(&schedule)

			resultLocked := database.Datasource.DB().
				Where("name", "RENEWAL_PUSH").
				Where("TIME(locked_at) = TIME(?)", timeNow).
				Where("status", 0).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				database.Datasource.DB().Save(&schedule)

				go func() {
					// populateRenewal()
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
				Where("name", "RENEWAL_PUSH").
				Where("TIME(publish_at) = TIME(?)", timeNow).
				Where("status", true).
				First(&schedule)

			resultLocked := database.Datasource.DB().
				Where("name", "RENEWAL_PUSH").
				Where("TIME(locked_at) = TIME(?)", timeNow).
				Where("status", false).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				database.Datasource.DB().Save(&schedule)

				go func() {
					// populateRenewal()
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
