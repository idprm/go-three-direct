package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DRRequest struct {
	MobileNo  string `json:"mobile_no"`
	ShortCode string `json:"short_code"`
	Message   string `json:"message"`
}

func DeliveryReport(c *gin.Context) {

	c.XML(http.StatusOK, gin.H{"status": "OK"})
}
