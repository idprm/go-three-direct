package app

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/wiliehidayat87/rmqp"
	"waki.mobi/go-yatta-h3i/src/config"
)

func StartApplication(cfg *config.Secret, db *sql.DB, rmpq rmqp.AMQP) *fiber.App {
	return mapUrls(cfg, db, rmpq)
}
