package cmd

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"sync"
	"time"

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

func moProccesor(wg *sync.WaitGroup, message []byte) {

	loggerMt := util.MakeLogger("mt", true)
	loggerNotif := util.MakeLogger("notif", true)
	loggerPb := util.MakeLogger("pb", true)

	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */
	transactionId := util.GenerateTransactionId()

	// init global variable
	var (
		content model.Content
	)

	// parsing string json
	var req dto.MORequest
	json.Unmarshal(message, &req)

	// get service by code
	service, _ := query.GetService(req.ShortCode)

	/**
	 * Query Content
	 */
	contFirstpush, _ := query.GetContent(service.ID, valFirstpush)

	contWelcome, _ := query.GetContent(service.ID, valWelcome)

	contWrongKey, _ := query.GetContent(service.ID, valErroyKey)

	// checking subscription
	var subscription model.Subscription
	activeSub := database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", req.MobileNo).First(&subscription)

	// split message param
	msg := strings.Split(req.Message, " ")
	// define array with index
	index0 := strings.ToUpper(msg[0])
	index1 := strings.ToUpper(msg[1])

	// split 5 character KEREN
	splitIndex1 := strings.ToUpper(string(msg[1][5:]))

	if activeSub.RowsAffected == 0 {

		database.Datasource.DB().Create(&model.Subscription{
			ServiceID: service.ID,
			Msisdn:    req.MobileNo,
			Keyword:   strings.ToUpper(req.Message),
			IpAddress: req.IpAddress,
			IsActive:  true,
		})
	}

	if activeSub.RowsAffected == 1 {
		subscription.IpAddress = req.IpAddress
		database.Datasource.DB().Save(&subscription)
	}

	/**
	 * Content keyword
	 */
	var contkeyword model.Keyword
	database.Datasource.DB().Where("name", splitIndex1).First(&contkeyword)

	var adnet model.Adnet
	database.Datasource.DB().Where("name", splitIndex1).First(&adnet)

	/**
	 * Error Keyword
	 */
	if (index0 != valReg && index0 != valUnreg) || index1 != service.Name {
		/**
		 * Push MT
		 */
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

	/**
	 * IF REG
	 */
	if index0 == valReg {
		/**
		 * Push MT
		 */
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

		/**
		 * Parsing XML Result
		 */
		resultFirstpush := util.EscapeChar(firstpushMt)
		responseXML := dto.Response{}
		xml.Unmarshal([]byte(resultFirstpush), &responseXML)
		submitedId := responseXML.Body.Param.SubmitedID
		resCode := responseXML.Body.Param.Code

		if resCode != 0 {
			subscription.SubmitedID = submitedId
			subscription.LatestStatus = statusFailed
			subscription.LatestSubject = smsFirstpush
			database.Datasource.DB().Save(&subscription)

			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					SubmitedID:    submitedId,
					Msisdn:        req.MobileNo,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        0,
					Subject:       smsFirstpush,
					Status:        statusFailed,
					StatusDetail:  util.ResponseCode(resCode),
					IpAddress:     req.IpAddress,
					Payload:       util.TrimByteToString(firstpushMt),
				},
			)

		} else {

			subscription.SubmitedID = submitedId
			subscription.LatestStatus = statusSuccess
			subscription.LatestSubject = smsFirstpush
			database.Datasource.DB().Save(&subscription)

			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					SubmitedID:    submitedId,
					Msisdn:        req.MobileNo,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        2200,
					Subject:       smsFirstpush,
					Status:        statusSuccess,
					StatusDetail:  util.ResponseCode(resCode),
					IpAddress:     req.IpAddress,
					Payload:       util.TrimByteToString(firstpushMt),
				},
			)

			/**
			 * Push Welcome
			 */
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

			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					SubmitedID:    submitedId,
					Msisdn:        req.MobileNo,
					Keyword:       strings.ToUpper(req.Message),
					Amount:        0,
					Subject:       smsWelcome,
					Status:        statusSuccess,
					IpAddress:     req.IpAddress,
					Payload:       util.TrimByteToString(welcomeMT),
				},
			)

			/**
			 * Postback
			 */
			postback, err := handler.Postback(service, req.MobileNo, transactionId)
			if err != nil {
				loggerPb.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"error":          err.Error(),
				}).Error()
			}
			loggerPb.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"payload":        util.TrimByteToString(postback),
			}).Info()

			/**
			 * Notif sub
			 */
			notifSub, err := handler.NotifSub(service, req.MobileNo, transactionId)
			if err != nil {
				loggerNotif.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"error":          err.Error(),
				}).Error()
			}
			loggerNotif.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"payload":        util.TrimByteToString(notifSub),
			}).Info()

		}

	}

	/**
	 * IF UNREG
	 */
	if index0 == valUnreg {

		unregMt, err := handler.MessageTerminated(service, content, req.MobileNo, transactionId)
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
			"payload":        util.TrimByteToString(unregMt),
		}).Info()

		resultUnreg := util.EscapeChar(unregMt)
		responseXML := dto.Response{}
		xml.Unmarshal([]byte(resultUnreg), &responseXML)
		submitedId := responseXML.Body.Param.SubmitedID
		statusDetail := responseXML.Body.Param.Text

		subscription.Keyword = strings.ToUpper(req.Message)
		subscription.SubmitedID = submitedId
		subscription.UnsubAt = time.Now()
		subscription.LatestSubject = smsUnsub
		subscription.IsRetry = false
		subscription.IsActive = false
		database.Datasource.DB().Save(&subscription)

		database.Datasource.DB().Create(
			&model.Transaction{
				TransactionID: transactionId,
				ServiceID:     service.ID,
				SubmitedID:    submitedId,
				Msisdn:        req.MobileNo,
				Keyword:       strings.ToUpper(req.Message),
				Amount:        0,
				Subject:       smsUnsub,
				Status:        "",
				StatusDetail:  statusDetail,
				IpAddress:     req.IpAddress,
				Payload:       util.TrimByteToString(unregMt),
			},
		)

		/**
		 * Notif Unsub
		 */
		notifUnsub, err := handler.NotifUnsub(service, req.MobileNo, transactionId)
		if err != nil {
			loggerNotif.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         req.MobileNo,
				"error":          err.Error(),
			}).Error()
		}
		loggerNotif.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         req.MobileNo,
			"payload":        util.TrimByteToString(notifUnsub),
		}).Info()

	}

	wg.Done()
}

func renewalProccesor(wg *sync.WaitGroup, message []byte) {
	loggerMt := util.MakeLogger("mt", true)

	transactionId := util.GenerateTransactionId()

	// parsing string json
	var sub model.Subscription
	json.Unmarshal(message, &sub)

	/**
	 * Query Service
	 */
	var service model.Service
	database.Datasource.DB().Where("id", sub.ServiceID).First(&service)

	/**
	 * Query Content
	 */
	var contRenewal model.Content
	database.Datasource.DB().Where("name", valRenewal).First(&contRenewal)

	var contInsuft model.Content
	database.Datasource.DB().Where("name", valInsuft).First(&contInsuft)

	renewalMt, err := handler.MessageTerminated(service, contRenewal, sub.Msisdn, transactionId)
	if err != nil {
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         sub.Msisdn,
			"error":          err.Error(),
		}).Error()
	}
	loggerMt.WithFields(logrus.Fields{
		"transaction_id": transactionId,
		"msisdn":         sub.Msisdn,
		"payload":        util.TrimByteToString(renewalMt),
	}).Info()

	resultRenewal := util.EscapeChar(renewalMt)
	responseXML := dto.Response{}
	xml.Unmarshal([]byte(resultRenewal), &responseXML)
	submitedId := responseXML.Body.Param.SubmitedID
	statusDetail := responseXML.Body.Param.Text

	var subscription model.Subscription
	database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", sub.Msisdn).First(&subscription)

	subscription.SubmitedID = submitedId
	subscription.LatestSubject = smsRenewal
	subscription.LatestStatus = ""
	database.Datasource.DB().Save(&subscription)

	database.Datasource.DB().Create(&model.Transaction{
		ServiceID:    sub.ServiceID,
		Msisdn:       sub.Msisdn,
		Keyword:      sub.Keyword,
		SubmitedID:   submitedId,
		Subject:      smsRenewal,
		Status:       "",
		StatusDetail: statusDetail,
		Payload:      util.TrimByteToString(renewalMt),
		IpAddress:    sub.IpAddress,
	})

	wg.Done()
}

func retryProccesor(wg *sync.WaitGroup, message []byte) {
	loggerMt := util.MakeLogger("mt", true)
	transactionId := util.GenerateTransactionId()

	// parsing string json
	var sub model.Subscription
	json.Unmarshal(message, &sub)

	/**
	 * Query Service
	 */
	var service model.Service
	database.Datasource.DB().Where("id", sub.ServiceID).First(&service)

	/**
	 * Query Content
	 */
	var contRenewal model.Content
	database.Datasource.DB().Where("name", valRenewal).First(&contRenewal)

	var contInsuft model.Content
	database.Datasource.DB().Where("name", valInsuft).First(&contInsuft)

	/**
	 * MT Retry
	 */
	retryMt, err := handler.MessageTerminated(service, contRenewal, sub.Msisdn, transactionId)
	if err != nil {
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         sub.Msisdn,
			"error":          err.Error(),
		}).Error("")
	}
	loggerMt.WithFields(logrus.Fields{
		"transaction_id": transactionId,
		"msisdn":         sub.Msisdn,
		"payload":        util.TrimByteToString(retryMt),
	}).Info("")

	resultRetry := util.EscapeChar(retryMt)
	responseXML := dto.Response{}
	xml.Unmarshal([]byte(resultRetry), &responseXML)
	submitedId := responseXML.Body.Param.SubmitedID

	var subscription model.Subscription
	database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", sub.Msisdn).First(&subscription)

	subscription.SubmitedID = submitedId
	database.Datasource.DB().Save(&subscription)

	wg.Done()
}

func drProccesor(wg *sync.WaitGroup, message []byte) {

	/**
	 * {"msisdn":"62895635121559","shortcode":"99879","status":"DELIVRD","message":"1601666764269215859","ip":"116.206.10.222"}
	 */

	// parsing string json
	var req dto.DRRequest
	json.Unmarshal(message, &req)

	// get service by code
	var service model.Service
	database.Datasource.DB().Where("code", req.ShortCode).First(&service)

	// get content
	var contRenewal model.Content
	database.Datasource.DB().Where("name", valRenewal).First(&contRenewal)

	var (
		status     string
		renewalDay int
	)

	// checking subscription
	var subscription model.Subscription
	activeSub := database.Datasource.DB().
		Where("service_id", service.ID).
		Where("msisdn", req.Msisdn).
		Where("submited_id", req.Message).
		First(&subscription)

	if activeSub.RowsAffected == 1 {
		if req.Status == "DELIVRD" {
			status = "SUCCESS"
			renewalDay = valRenewalSuccessAt
			subscription.ChargeAt = time.Now()
			subscription.IsRetry = false

		} else {
			status = "FAILED"
			renewalDay = valRenewalFailedAt
			subscription.RetryAt = time.Now()
			subscription.IsRetry = true
		}

		subscription.LatestSubject = valRenewal
		subscription.LatestStatus = status
		subscription.Success = subscription.Success + 1
		subscription.RenewalAt = time.Now().AddDate(0, 0, renewalDay)
		database.Datasource.DB().Save(&subscription)
	}

	var transaction model.Transaction
	getTransaction := database.Datasource.DB().
		Where("service_id", service.ID).
		Where("msisdn", req.Msisdn).
		Where("submited_id", req.Message).First(&transaction)

	if getTransaction.RowsAffected == 1 {
		transaction.Status = status
		transaction.Amount = 0
		transaction.Subject = valRenewal
		database.Datasource.DB().Save(&transaction)
	}

	wg.Done()
}
