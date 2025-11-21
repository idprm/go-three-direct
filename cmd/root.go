package cmd

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/wiliehidayat87/rmqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	APP_HOST  string = getEnv("APP_HOST")
	APP_PORT  string = getEnv("APP_PORT")
	APP_TZ    string = getEnv("APP_TZ")
	APP_URL   string = getEnv("APP_URL")
	URI_MYSQL string = getEnv("URI_MYSQL")
	URI_REDIS string = getEnv("URI_REDIS")
	URI_AMQP  string = getEnv("URI_AMQP")
	RMQ_HOST  string = getEnv("RMQ_HOST")
	RMQ_USER  string = getEnv("RMQ_USER")
	RMQ_PASS  string = getEnv("RMQ_PASS")
	RMQ_PORT  string = getEnv("RMQ_PORT")
	RMQ_URL   string = getEnv("RMQ_URL")
	LOG_PATH  string = getEnv("LOG_PATH")
)

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	/*
	 * Listener service
	 */
	rootCmd.AddCommand(listenerCmd)

	/**
	 *  Consumer service
	 */
	rootCmd.AddCommand(consumerMOCmd)
	rootCmd.AddCommand(consumerDRCmd)
	rootCmd.AddCommand(consumerRenewalCmd)
	rootCmd.AddCommand(consumerRetryCmd)
	rootCmd.AddCommand(consumerNotif)

	/**
	 *  Publisher service
	 */
	rootCmd.AddCommand(publisherRenewalCmd)
	rootCmd.AddCommand(publisherRetryCmd)

}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return value
}

func connectSQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", URI_MYSQL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(URI_MYSQL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Connect to rabbitmq
func connectRabbitMq() rmqp.AMQP {
	var rb rmqp.AMQP
	port, _ := strconv.Atoi(RMQ_PORT)
	rb.SetAmqpURL(RMQ_HOST, port, RMQ_USER, RMQ_PASS)
	rb.SetUpConnectionAmqp()
	return rb
}

// Connect to redis
func connectRedis() (*redis.Client, error) {
	opts, err := redis.ParseURL(URI_REDIS)
	if err != nil {
		return nil, err
	}
	return redis.NewClient(opts), nil
}
