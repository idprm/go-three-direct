package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

func DeliveryReport(c *gin.Context) {
	/**
	 * {"msisdn":"62895330590144","shortcode":"998791","status":"DELIVRD","message":"1601666588632810494","ip":"116.206.10.222"}
	 */
	var req dto.DRRequest

	/**
	 * Body Parsing
	 */
	if err := c.ShouldBindQuery(&req); err != nil {
		c.XML(http.StatusBadRequest, gin.H{
			"error":   true,
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	json, _ := json.Marshal(req)

	queue.Rabbit.IntegratePublish(
		config.ViperEnv("RMQ_DREXCHANGE"),
		config.ViperEnv("RMQ_DRQUEUE"),
		config.ViperEnv("RMQ_DRDATATYPE"),
		"",
		string(json),
	)

	c.XML(http.StatusOK, gin.H{
		"error":   false,
		"code":    http.StatusOK,
		"message": "Success",
	})
	return
}
