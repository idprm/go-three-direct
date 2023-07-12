package rabbitmq

import (
	"github.com/wiliehidayat87/rmqp"
	"waki.mobi/go-yatta-h3i/src/config"
)

func InitQueue(cfg *config.Secret) rmqp.AMQP {
	var rb rmqp.AMQP

	rb.SetAmqpURL(cfg.Rmq.Host, cfg.Rmq.Port, cfg.Rmq.User, cfg.Rmq.Pass)

	rb.SetUpConnectionAmqp()
	return rb
}
