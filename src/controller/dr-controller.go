package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

type DRRequest struct {
	Msisdn    string `query:"msisdn" json:"msisdn"`
	ShortCode string `query:"shortcode" json:"shortcode"`
	Status    string `query:"status" json:"status"`
	Message   string `query:"message" json:"message"`
	IpAddress string `query:"ip" json:"ip"`
}

func DeliveryReport(c *gin.Context) {
	/**
	 * {"msisdn":"62895330590144","shortcode":"998791","status":"DELIVRD","message":"1601666588632810494","ip":"116.206.10.222"}
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
		config.ViperEnv("RMQ_DREXCHANGE"),
		true,
		config.ViperEnv("RMQ_DRQUEUE"),
	)

	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
