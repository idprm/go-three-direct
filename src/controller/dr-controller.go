package controller

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

func DeliveryReport(c *fiber.Ctx) error {
	/**
	 * {"msisdn":"62895330590144","shortcode":"998791","status":"DELIVRD","message":"1601666588632810494","ip":"116.206.10.222"}
	 */

	/**
	 * Query Parser
	 */
	req := new(dto.DRRequest)

	if err := c.QueryParser(req); err != nil {
		return err
	}

	json, _ := json.Marshal(req)

	queue.Rabbit.IntegratePublish(
		"E_DR",
		"Q_DR",
		"application/json",
		"",
		string(json),
	)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"code":    fiber.StatusOK,
		"message": "Success",
	})
}
