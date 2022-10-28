package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/handler"
	"waki.mobi/go-yatta-h3i/src/pkg/query"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

const (
	valReg          = "REG"
	valUnreg        = "UNREG"
	valWelcome      = "WELCOME"
	valRegistration = "REGISTRATION"
	valConfirmation = "CONFIRMATION"
	valFirstpush    = "FIRSTPUSH"
	valRenewal      = "RENEWAL"
	valUnsub        = "UNSUB"
	valInsuft       = "INSUFT"
	valErroyKey     = "ERROR_KEYWORD"
	valFailed       = "FAILED"
	valReminder     = "REMINDER"
	valIsActive     = "IS_ACTIVE"
	valPurge        = "PURGE"

	smsFirstpush = "MT_FIRSTPUSH"
	smsWelcome   = "MT_WELCOME"
	smsRenewal   = "MT_RENEWAL"
	smsInsuff    = "MT_INSUFFICIENT"
	smsUnsub     = "MT_UNSUB"

	statusFailed  = "FAILED"
	statusSuccess = "SUCCESS"

	valRenewalSuccessAt = 2
	valRenewalFailedAt  = 1
	valReminderAt       = 119
	valPurgeAt          = 120
)

func TestMO(c *fiber.Ctx) error {
	transactionId := util.GenerateTransactionId()
	loggerMt := util.MakeLogger("mt", true)

	/**
	 * Query Parser
	 */
	req := new(dto.MORequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	// get service by code
	service, _ := query.GetServiceByCode(req.ShortCode)

	// split message param
	msg := strings.Split(req.Message, " ")
	// define array with index
	index0 := strings.ToUpper(msg[0])
	index1 := strings.ToUpper(msg[1])

	// split 5 character KEREN[SPLIT]
	// splitIndex1 := strings.ToUpper(string(msg[1][5:]))

	/**
	 * Query Content
	 */
	contFirstpush, _ := query.GetContent(service.ID, valFirstpush)

	contWelcome, _ := query.GetContent(service.ID, valWelcome)

	contUnsub, _ := query.GetContent(service.ID, valUnsub)

	contWrongKey, _ := query.GetContent(service.ID, valErroyKey)

	/**
	 * FILTER BY MESSAGE
	 */
	if index0 == valReg && index1 == service.Name {
		firstpushMt, err := handler.MessageTerminated(service, contFirstpush, req.MobileNo, transactionId)
		if err != nil {
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"error":          err.Error(),
			}).Error()
		}
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         req.MobileNo,
			"payload":        util.TrimByteToString(firstpushMt),
		}).Info()

		welcomeMT, err := handler.MessageTerminated(service, contWelcome, req.MobileNo, transactionId)
		if err != nil {
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"error":          err.Error(),
			}).Error()
		}
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         req.MobileNo,
			"payload":        util.TrimByteToString(welcomeMT),
		}).Info()

	} else if index0 == valUnreg {
		unsubMt, err := handler.MessageTerminated(service, contUnsub, req.MobileNo, transactionId)
		if err != nil {
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"error":          err.Error(),
			}).Error()
		}
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         req.MobileNo,
			"payload":        util.TrimByteToString(unsubMt),
		}).Info()
	} else {

		wrongKeywordMt, err := handler.MessageTerminated(service, contWrongKey, req.MobileNo, transactionId)
		if err != nil {
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"error":          err.Error(),
			}).Error()
		}
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         req.MobileNo,
			"payload":        util.TrimByteToString(wrongKeywordMt),
		}).Info()
	}

	return c.XML(dto.ResponseXML{
		Status: "OK",
	})
}

func TestDR(c *fiber.Ctx) error {

	/**
	 * Query Parser
	 */
	req := new(dto.DRRequest)
	if err := c.QueryParser(req); err != nil {
		return err
	}

	return c.XML(dto.ResponseXML{
		Status: "OK",
	})
}
