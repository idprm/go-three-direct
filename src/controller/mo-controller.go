package controller

import (
	"database/sql"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/query"
	"waki.mobi/go-yatta-h3i/src/pkg/queue"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

type MOHandler struct {
	db  *sql.DB
	gdb *gorm.DB
}

func NewHandlerMO(db *sql.DB, gdb *gorm.DB) *MOHandler {
	return &MOHandler{
		db:  db,
		gdb: gdb,
	}
}

func (h *MOHandler) MessageOriginated(c *fiber.Ctx) error {
	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */
	loggerMo := util.MakeLogger("mo", true)

	/**
	 * Query Parser
	 */
	req := new(dto.MORequest)

	req.IpAddress = c.IP()

	if err := c.QueryParser(req); err != nil {
		return err
	}

	loggerMo.WithFields(logrus.Fields{
		"request": req,
	}).Info()

	blacklistRepo := query.NewBlacklistRepository(h.db)

	/**
	 * If MSISDN is blacklist
	 */
	count, _ := blacklistRepo.GetCountBlacklist(req.MobileNo)

	if count == 0 {
		json, _ := json.Marshal(req)

		queue.Rabbit.IntegratePublish(
			"E_MO",
			"Q_MO",
			"application/json",
			"",
			string(json),
		)
	}

	return c.XML(dto.ResponseXML{
		Status: "OK",
	})
}
