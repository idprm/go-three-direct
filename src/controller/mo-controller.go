package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
)

func MessageOriginated(c *gin.Context) {
	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */

	var req dto.MORequest
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

	msg := strings.Split(req.Message, " ")
	log.Println(msg[0])
	log.Println(msg[1])

	// json, _ := json.Marshal(req)

	// if req.MobileNo != "" {

	// 	queue.Rabbit.IntegratePublish(
	// 		config.ViperEnv("RMQ_MOEXCHANGE"),
	// 		config.ViperEnv("RMQ_MOQUEUE"),
	// 		config.ViperEnv("RMQ_MODATATYPE"),
	// 		"",
	// 		string(json),
	// 	)
	// }

	// push to rabbitmq
	c.XML(http.StatusOK, gin.H{
		"error":   false,
		"code":    http.StatusOK,
		"message": "Success",
	})
	return
}
