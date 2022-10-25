package queue

import (
	"strconv"

	"github.com/wiliehidayat87/rmqp"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
)

var Rabbit rmqp.AMQP

func SetupQueue() {

	port, _ := strconv.Atoi(config.ViperEnv("RMQ_PORT"))

	Rabbit.SetAmqpURL(
		config.ViperEnv("RMQ_HOST"),
		port,
		config.ViperEnv("RMQ_USER"),
		config.ViperEnv("RMQ_PASS"),
	)

	Rabbit.SetUpConnectionAmqp()
}
