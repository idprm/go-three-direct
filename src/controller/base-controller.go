package controller

import (
	"github.com/gofiber/fiber/v2"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func ReportMO(c *fiber.Ctx) error {

	var transactions []model.Transaction
	database.Datasource.DB().Select("count(1) as subject, keyword, adnet, status, DATE(created_at) as created_at").
		Where("DATE(created_at) BETWEEN DATE_SUB(NOW(), INTERVAL 30 DAY) AND DATE(NOW())").
		Where("subject", "MT_FIRSTPUSH").
		Group("DATE(created_at), adnet").
		Order("DATE(created_at) DESC").Find(&transactions)

	return c.Render("report/mo", fiber.Map{
		"transactions": transactions,
	})
}

func ReportRenewal(c *fiber.Ctx) error {
	var transactions []model.Transaction
	database.Datasource.DB().Select("DATE(created_at) as created_at, status, count(1) as subject").
		Where("DATE(created_at) BETWEEN DATE_SUB(NOW(), INTERVAL 30 DAY) AND DATE(NOW())").
		Where("subject", "MT_RENEWAL").
		Group("DATE(created_at), status").
		Order("DATE(created_at) DESC").Find(&transactions)

	return c.Render("report/renewal", fiber.Map{
		"transactions": transactions,
	})
}

func ReportFirstpush(c *fiber.Ctx) error {
	var transactions []model.Transaction
	database.Datasource.DB().Select("DATE(created_at) as created_at, status, count(1) as subject").
		Where("DATE(created_at) BETWEEN DATE_SUB(NOW(), INTERVAL 30 DAY) AND DATE(NOW())").
		Where("subject", "MT_FIRSTPUSH").
		Group("DATE(created_at), status").
		Order("DATE(created_at) DESC").Find(&transactions)

	return c.Render("report/firstpush", fiber.Map{
		"transactions": transactions,
	})
}
