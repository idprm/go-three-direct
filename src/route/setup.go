package route

import (
	"github.com/gin-gonic/gin"
	"waki.mobi/go-yatta-h3i/src/controller"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/moh3i", controller.MessageOriginated)
	router.GET("/drh3i", controller.DeliveryReport)

	return router
}
