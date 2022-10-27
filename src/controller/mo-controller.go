package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

func MessageOriginated(c *fiber.Ctx) error {
	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */
	loggerMo := util.MakeLogger("mo", true)

	/**
	 * Query Parser
	 */
	req := new(dto.MORequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	loggerMo.WithFields(logrus.Fields{
		"request": req,
	}).Info()

	// splitIndex1 := strings.ToUpper(string(msg[1][5:]))

	json, _ := json.Marshal(req)

	queue.Rabbit.IntegratePublish(
		"E_MO",
		"Q_MO",
		"application/json",
		"",
		string(json),
	)

	return c.XML(dto.ResponseXML{
		Status: "OK",
	})
}
