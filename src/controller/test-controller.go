package controller

import (
	"encoding/xml"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/handler"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
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
	smsIsActive  = "MT_ISACTIVE"
	smsPurge     = "MT_PURGE"
	smsRenewal   = "MT_RENEWAL"
	smsInsuff    = "MT_INSUFFICIENT"
	smsUnsub     = "MT_UNSUB"
	smsWrongKey  = "MT_WRONGKEY"

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

	contIsActive, _ := query.GetContent(service.ID, valIsActive)

	contUnsub, _ := query.GetContent(service.ID, valUnsub)

	contPurge, _ := query.GetContent(service.ID, valPurge)

	contWrongKey, _ := query.GetContent(service.ID, valErroyKey)

	/**
	 * FILTER BY MESSAGE
	 */
	if index0 == valReg && index1 == service.Name {

		var subscription model.Subscription
		existSub := database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", req.MobileNo).Where("is_active", true).First(&subscription)

		// IF SUB NOT EXIST
		if existSub.RowsAffected == 0 {

			var (
				labelStatus  string
				dayRenewal   int
				purgeAt      time.Time
				chargeAt     time.Time
				isRetry      bool
				chargeAmount float64
			)

			// sent mt_firstpush
			firstpushMt, err := handler.MessageTerminated(service, contFirstpush, req.MobileNo, transactionId)
			if err != nil {
				loggerMt.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"error":          err.Error(),
				}).Error(smsFirstpush)
			}
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"payload":        util.TrimByteToString(firstpushMt),
			}).Info(smsFirstpush)

			resultFirstpush := util.EscapeChar(firstpushMt)
			resXML := dto.Response{}
			xml.Unmarshal([]byte(resultFirstpush), &resXML)
			submitedId := resXML.Body.Param.SubmitedID
			statusCode := resXML.Body.Param.Code

			// if status code 0 = success
			if statusCode == 0 {
				labelStatus = "SUCCESS"
				dayRenewal = service.Day
				purgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
				chargeAt = time.Now()
				chargeAmount = service.Charge
				isRetry = false
			} else {
				labelStatus = "FAILED"
				dayRenewal = 1
				purgeAt = time.Time{}
				chargeAt = time.Time{}
				chargeAmount = 0
				isRetry = true
			}

			// Insert to subscription
			database.Datasource.DB().Create(
				&model.Subscription{
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					Keyword:       strings.ToUpper(req.Message),
					LatestSubject: smsFirstpush,
					LatestStatus:  labelStatus,
					Amount:        chargeAmount,
					RenewalAt:     time.Now().AddDate(0, 0, dayRenewal),
					PurgeAt:       purgeAt,
					ChargeAt:      chargeAt,
					Success:       1,
					IpAddress:     "",
					IsRetry:       isRetry,
					IsActive:      true,
				},
			)

			// Insert to Transaction
			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					SubmitedID:    submitedId,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        0,
					Status:        labelStatus,
					StatusDetail:  util.ResponseStatusCode(statusCode),
					Subject:       smsFirstpush,
					IpAddress:     "",
					Payload:       util.TrimByteToString(firstpushMt),
				},
			)

			// sent mt_welcome
			welcomeMT, err := handler.MessageTerminated(service, contWelcome, req.MobileNo, transactionId)
			if err != nil {
				loggerMt.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"error":          err.Error(),
				}).Error(smsWelcome)
			}
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"payload":        util.TrimByteToString(welcomeMT),
			}).Info(smsWelcome)

			resultWelcome := util.EscapeChar(welcomeMT)
			res1XML := dto.Response{}
			xml.Unmarshal([]byte(resultWelcome), &res1XML)
			submitedIdwelcome := resXML.Body.Param.SubmitedID
			statusCodewelcome := resXML.Body.Param.Code

			// Insert to Transaction
			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					SubmitedID:    submitedIdwelcome,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        0,
					Status:        "",
					StatusDetail:  util.ResponseStatusCode(statusCodewelcome),
					Subject:       smsWelcome,
					IpAddress:     "",
					Payload:       util.TrimByteToString(welcomeMT),
				},
			)
		}

		// IF SUB EXIST
		if existSub.RowsAffected == 1 {
			// sent mt_is_active
			isActiveMT, err := handler.MessageTerminated(service, contIsActive, req.MobileNo, transactionId)
			if err != nil {
				loggerMt.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"error":          err.Error(),
				}).Error(smsIsActive)
			}
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"payload":        util.TrimByteToString(isActiveMT),
			}).Info(smsIsActive)

			resultIsActive := util.EscapeChar(isActiveMT)
			resXML := dto.Response{}
			xml.Unmarshal([]byte(resultIsActive), &resXML)
			submitedId := resXML.Body.Param.SubmitedID
			statusCode := resXML.Body.Param.Code

			// Insert to Transaction
			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					SubmitedID:    submitedId,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        0,
					Status:        "",
					StatusDetail:  util.ResponseStatusCode(statusCode),
					Subject:       smsIsActive,
					IpAddress:     "",
					Payload:       util.TrimByteToString(isActiveMT),
				},
			)
		}

	} else if index0 == valUnreg {

		var subscription model.Subscription
		existSub := database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", req.MobileNo).Where("is_active", true).First(&subscription)

		// IF SUB EXIST
		if existSub.RowsAffected == 1 {

			// sent mt_unsub
			unsubMT, err := handler.MessageTerminated(service, contUnsub, req.MobileNo, transactionId)
			if err != nil {
				loggerMt.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"error":          err.Error(),
				}).Error(smsUnsub)
			}
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"payload":        util.TrimByteToString(unsubMT),
			}).Info(smsUnsub)

			resultUnsub := util.EscapeChar(unsubMT)
			resXML := dto.Response{}
			xml.Unmarshal([]byte(resultUnsub), &resXML)
			submitedId := resXML.Body.Param.SubmitedID
			statusCode := resXML.Body.Param.Code

			// Insert to Transaction
			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					SubmitedID:    submitedId,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        0,
					Status:        "",
					StatusDetail:  util.ResponseStatusCode(statusCode),
					Subject:       smsUnsub,
					IpAddress:     "",
					Payload:       util.TrimByteToString(unsubMT),
				},
			)
		}

		// IF SUB NOT EXIST
		if existSub.RowsAffected == 0 {
			// sent mt_purge
			purgeMT, err := handler.MessageTerminated(service, contPurge, req.MobileNo, transactionId)
			if err != nil {
				loggerMt.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"error":          err.Error(),
				}).Error(smsPurge)
			}
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"payload":        util.TrimByteToString(purgeMT),
			}).Info(smsPurge)

			resultPurge := util.EscapeChar(purgeMT)
			resXML := dto.Response{}
			xml.Unmarshal([]byte(resultPurge), &resXML)
			submitedId := resXML.Body.Param.SubmitedID
			statusCode := resXML.Body.Param.Code

			// Insert to Transaction
			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					SubmitedID:    submitedId,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        0,
					Status:        "",
					StatusDetail:  util.ResponseStatusCode(statusCode),
					Subject:       smsPurge,
					IpAddress:     "",
					Payload:       util.TrimByteToString(purgeMT),
				},
			)
		}

	} else {
		// sent mt_wrongkey
		wrongKeywordMt, err := handler.MessageTerminated(service, contWrongKey, req.MobileNo, transactionId)
		if err != nil {
			loggerMt.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"error":          err.Error(),
			}).Error(smsWrongKey)
		}
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         req.MobileNo,
			"payload":        util.TrimByteToString(wrongKeywordMt),
		}).Info(smsWrongKey)

		resultPurge := util.EscapeChar(wrongKeywordMt)
		resXML := dto.Response{}
		xml.Unmarshal([]byte(resultPurge), &resXML)
		submitedId := resXML.Body.Param.SubmitedID
		statusCode := resXML.Body.Param.Code

		// Insert to Transaction
		database.Datasource.DB().Create(
			&model.Transaction{
				TransactionID: transactionId,
				ServiceID:     service.ID,
				Msisdn:        req.MobileNo,
				SubmitedID:    submitedId,
				Keyword:       strings.ToUpper(req.Message),
				Amount:        0,
				Status:        "",
				StatusDetail:  util.ResponseStatusCode(statusCode),
				Subject:       smsWrongKey,
				IpAddress:     "",
				Payload:       util.TrimByteToString(wrongKeywordMt),
			},
		)
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
