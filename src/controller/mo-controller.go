package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

type MORequest struct {
	MobileNo  string `form:"mobile_no" json:"mobile_no"`
	ShortCode string `form:"short_code" json:"short_code"`
	Message   string `form:"message" json:"message"`
	IpAddress string `form:"ip" json:"ip"`
}

func MessageOriginated(c *gin.Context) {
	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */

	var req MORequest

	req.MobileNo = c.Query("mobile_no")
	req.ShortCode = c.Query("short_code")
	req.Message = c.Query("message")
	req.IpAddress = c.Query("ip")

	json, _ := json.Marshal(req)
	isPublished := queue.Rabbit.IntegratePublish(
		config.ViperEnv("RMQ_MOEXCHANGE"),
		config.ViperEnv("RMQ_MOQUEUE"),
		config.ViperEnv("RMQ_MODATATYPE"),
		"",
		string(json),
	)

	if isPublished {
		fmt.Println("[v] Published")
	} else {
		fmt.Println("[v] Failed Published")
	}

	// push to rabbitmq
	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
