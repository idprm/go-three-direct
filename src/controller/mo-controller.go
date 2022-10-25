package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

type MORequest struct {
	MobileNo  string `query:"mobile_no" json:"mobile_no"`
	ShortCode string `query:"short_code" json:"short_code"`
	Message   string `query:"message" json:"message"`
	IpAddress string `query:"ip" json:"ip"`
}

func MessageOriginated(c *gin.Context) {

	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */

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
		config.ViperEnv("RMQ_MOEXCHANGE"),
		true,
		config.ViperEnv("RMQ_MOQUEUE"),
	)

	// push to rabbitmq
	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
