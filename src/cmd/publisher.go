package cmd

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/spf13/cobra"
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/database/mysql/db"
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
		 * LOAD CONFIG
		 */
		cfg, err := config.LoadSecret("secret.yaml")
		if err != nil {
			panic(err)
		}

		/**
		 * SETUP MYSQL
		 */
		sdb := db.InitDB(cfg)
		gdb := db.InitGormDB(cfg)

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
			resultPublish := gdb.
				Where("name", "RENEWAL_PUSH").
				Where("TIME(publish_at) = TIME(?)", timeNow).
				Where("status", true).
				First(&schedule)

			resultLocked := gdb.
				Where("name", "RENEWAL_PUSH").
				Where("TIME(un_locked_at) = TIME(?)", timeNow).
				Where("status", false).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				gdb.Save(&schedule)

				go func() {
					populateRenewal(sdb)
				}()

			}

			if resultLocked.RowsAffected == 1 {
				schedule.Status = true
				gdb.Save(&schedule)
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
		 * LOAD CONFIG
		 */
		cfg, err := config.LoadSecret("secret.yaml")
		if err != nil {
			panic(err)
		}

		/**
		 * SETUP MYSQL
		 */
		sdb := db.InitDB(cfg)
		gdb := db.InitGormDB(cfg)

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
			resultPublish := gdb.
				Where("name", "RETRY_PUSH").
				Where("TIME(publish_at) = TIME(?)", timeNow).
				Where("status", true).
				First(&schedule)

			resultLocked := gdb.
				Where("name", "RETRY_PUSH").
				Where("TIME(un_locked_at) = TIME(?)", timeNow).
				Where("status", false).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				gdb.Save(&schedule)

				go func() {
					populateRetry(sdb)
				}()

			}

			if resultLocked.RowsAffected == 1 {
				schedule.Status = true
				gdb.Save(&schedule)
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
		 * LOAD CONFIG
		 */
		cfg, err := config.LoadSecret("secret.yaml")
		if err != nil {
			panic(err)
		}

		/**
		 * SETUP MYSQL
		 */
		sdb := db.InitDB(cfg)
		gdb := db.InitGormDB(cfg)

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
			resultPublish := gdb.
				Where("name", "PURGE_PUSH").
				Where("TIME(publish_at) = TIME(?)", timeNow).
				Where("status", true).
				First(&schedule)

			resultLocked := gdb.
				Where("name", "PURGE_PUSH").
				Where("TIME(un_locked_at) = TIME(?)", timeNow).
				Where("status", false).
				First(&schedule)

			if resultPublish.RowsAffected == 1 {
				schedule.Status = false
				gdb.Save(&schedule)

				go func() {
					populatePurge(sdb)
				}()

			}

			if resultLocked.RowsAffected == 1 {
				schedule.Status = true
				gdb.Save(&schedule)
			}

			time.Sleep(timeDuration * time.Minute)

		}

	},
}

func populateRenewal(sdb *sql.DB) {

	populateRepo := query.NewPopulateRepository(sdb)

	subs, _ := populateRepo.GetDataPopulate("RENEWAL")

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

func populateRetry(sdb *sql.DB) {

	populateRepo := query.NewPopulateRepository(sdb)

	subs, _ := populateRepo.GetDataPopulate("RETRY")

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

func populatePurge(sdb *sql.DB) {
	populateRepo := query.NewPopulateRepository(sdb)

	subs, _ := populateRepo.GetDataPopulate("PURGE")

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
