package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

func DeliveryReport(c *fiber.Ctx) error {
	/**
	 * {"msisdn":"62895330590144","shortcode":"998791","status":"DELIVRD","message":"1601666588632810494","ip":"116.206.10.222"}
	 */
	loggerDr := util.MakeLogger("dr", true)

	/**
	 * Query Parser
	 */
	req := new(dto.DRRequest)

	if err := c.QueryParser(req); err != nil {
		return err
	}

	loggerDr.WithFields(logrus.Fields{
		"request": req,
	}).Info()

	json, _ := json.Marshal(req)

	queue.Rabbit.IntegratePublish(
		"E_DR",
		"Q_DR",
		"application/json",
		"",
		string(json),
	)

	return c.XML(dto.ResponseXML{
		Status: "OK",
	})
}
