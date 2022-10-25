package controller

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
)

func MessageOriginated(c *fiber.Ctx) error {
	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */
	/**
	 * Query Parser
	 */
	req := new(dto.MORequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	msg := strings.Split(req.Message, " ")
	log.Println(msg[0])
	log.Println(msg[1])

	json, _ := json.Marshal(req)

	queue.Rabbit.IntegratePublish(
		"E_MO",
		"Q_MO",
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
