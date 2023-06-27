package controller

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

type IncomingHandler struct {
	db  *sql.DB
	gdb *gorm.DB
}

func NewIncomingHandler(db *sql.DB, gdb *gorm.DB) *IncomingHandler {
	return &IncomingHandler{
		db:  db,
		gdb: gdb,
	}
}

func (h *IncomingHandler) ReportMO(c *fiber.Ctx) error {

	var transactions []model.Transaction
	h.gdb.Select("count(1) as subject, keyword, adnet, status, DATE(created_at) as created_at").
		Where("DATE(created_at) BETWEEN DATE_SUB(NOW(), INTERVAL 30 DAY) AND DATE(NOW())").
		Where("subject", "MT_FIRSTPUSH").
		Group("DATE(created_at), adnet").
		Order("DATE(created_at) DESC").Find(&transactions)

	return c.Render("report/mo", fiber.Map{
		"transactions": transactions,
	})
}

func (h *IncomingHandler) ReportRenewal(c *fiber.Ctx) error {
	var transactions []model.Transaction
	h.gdb.Select("DATE(created_at) as created_at, status, count(1) as subject").
		Where("DATE(created_at) BETWEEN DATE_SUB(NOW(), INTERVAL 30 DAY) AND DATE(NOW())").
		Where("subject", "MT_RENEWAL").
		Group("DATE(created_at), status").
		Order("DATE(created_at) DESC").Find(&transactions)

	return c.Render("report/renewal", fiber.Map{
		"transactions": transactions,
	})
}

func (h *IncomingHandler) ReportFirstpush(c *fiber.Ctx) error {
	var transactions []model.Transaction
	h.gdb.Select("DATE(created_at) as created_at, status, count(1) as subject").
		Where("DATE(created_at) BETWEEN DATE_SUB(NOW(), INTERVAL 30 DAY) AND DATE(NOW())").
		Where("subject", "MT_FIRSTPUSH").
		Group("DATE(created_at), status").
		Order("DATE(created_at) DESC").Find(&transactions)

	return c.Render("report/firstpush", fiber.Map{
		"transactions": transactions,
	})
}
