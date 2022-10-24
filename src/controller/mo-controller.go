package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MORequest struct {
	MobileNo  string `json:"mobile_no"`
	ShortCode string `json:"short_code"`
	Message   string `json:"message"`
}

func MessageOriginated(c *gin.Context) {
	// http://35.247.131.49/moh3i?mobile_no=6289501845333&short_code=99879&message=reg+keren

	// push to rabbitmq
	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
