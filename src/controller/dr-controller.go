package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

type DRRequest struct {
	Msisdn    string `form:"msisdn" json:"msisdn"`
	ShortCode string `form:"shortcode" json:"shortcode"`
	Status    string `form:"status" json:"status"`
	Message   string `form:"message" json:"message"`
	IpAddress string `form:"ip" json:"ip"`
}

func DeliveryReport(c *gin.Context) {
	/**
	 * {"msisdn":"62895330590144","shortcode":"998791","status":"DELIVRD","message":"1601666588632810494","ip":"116.206.10.222"}
	 */

	var req DRRequest
	if c.ShouldBindQuery(&req) == nil {
		json, _ := json.Marshal(req)

		go func() {
			queue.Rabbit.IntegratePublish(
				config.ViperEnv("RMQ_DREXCHANGE"),
				config.ViperEnv("RMQ_DRQUEUE"),
				config.ViperEnv("RMQ_DRDATATYPE"),
				"",
				string(json),
			)
		}()

	}

	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
