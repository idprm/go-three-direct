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
	valWelcome   = "WELCOME"
	valFirstpush = "FIRSTPUSH"
	valUnsub     = "UNSUB"
	valInsuft    = "INSUFT"
	valErroyKey  = "ERROR_KEYWORD"
	valIsActive  = "IS_ACTIVE"
	valPurge     = "PURGE"
	valUnknown   = "UNKNOWN_KEYWORD"

	smsFirstpush = "MT_FIRSTPUSH"
	smsWelcome   = "MT_WELCOME"
	smsIsActive  = "MT_ISACTIVE"
	smsPurge     = "MT_PURGE"
	smsRenewal   = "MT_RENEWAL"
	smsInsuff    = "MT_INSUFFICIENT"
	smsUnsub     = "MT_UNSUB"
	smsWrongKey  = "MT_WRONGKEY"
)

func moProccesor(wg *sync.WaitGroup, message []byte) {

	/**
	 * Sample Request
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */
	loggerMt := util.MakeLogger("mt", true)
	loggerNotif := util.MakeLogger("notif", true)
	loggerPb := util.MakeLogger("postback", true)

	transactionId := util.GenerateTransactionId()

	// parsing string json
	var req dto.MORequest
	json.Unmarshal(message, &req)

	/**
	 * Query content
	 */
	contentUnknown, _ := query.GetContent(2, valUnknown)

	// get service by name
	service, _ := query.GetServiceByName(util.FilterMessage(strings.ToUpper(req.Message)))

	if (service.Name != "KEREN" && service.Name != "GM") || service.ID == 0 {
		unknownKeywordMt, err := handler.MessageTerminatedUnknown(contentUnknown, req.MobileNo, transactionId)
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
			"payload":        util.TrimByteToString(unknownKeywordMt),
		}).Info(smsWrongKey)

		resultWrongkey := util.EscapeChar(unknownKeywordMt)
		resXML := dto.Response{}
		xml.Unmarshal([]byte(resultWrongkey), &resXML)
		submitedId := resXML.Body.SubmitedID
		statusCode := resXML.Body.Code
		statusText := resXML.Body.Text

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
				StatusCode:    statusCode,
				StatusDetail:  statusText,
				Subject:       smsWrongKey,
				IpAddress:     "",
				Payload:       util.TrimByteToString(unknownKeywordMt),
			},
		)
	} else {
		/**
		 * Query Content
		 */
		contFirstpush, _ := query.GetContent(service.ID, valFirstpush)

		contWelcome, _ := query.GetContent(service.ID, valWelcome)

		contInsuff, _ := query.GetContent(service.ID, valInsuft)

		contIsActive, _ := query.GetContent(service.ID, valIsActive)

		contUnsub, _ := query.GetContent(service.ID, valUnsub)

		contPurge, _ := query.GetContent(service.ID, valPurge)

		contWrongKey, _ := query.GetContent(service.ID, valErroyKey)

		var subHasActive model.Subscription
		existSub := database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", req.MobileNo).Where("is_active", true).First(&subHasActive)

		var subInActive model.Subscription
		nonActiveSub := database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", req.MobileNo).Where("is_active", false).First(&subInActive)

		adn := util.KeywordDefine(strings.ToUpper(req.Message))

		var adnet model.Adnet
		database.Datasource.DB().Where("name", adn).First(&adnet)

		/**
		 * IF SUB IS EXIST AND IS_ACTIVE = true
		 */
		if existSub.RowsAffected == 1 && util.FilterReg(req.Message) {
			subHasActive.Keyword = strings.ToUpper(req.Message)
			subHasActive.Adnet = adnet.Value
			subHasActive.IpAddress = req.IpAddress
			database.Datasource.DB().Save(&subHasActive)

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
			submitedId := resXML.Body.SubmitedID
			statusCode := resXML.Body.Code
			statusText := resXML.Body.Text

			// Insert to Transaction
			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					SubmitedID:    submitedId,
					Keyword:       strings.ToUpper(req.Message),
					Adnet:         adnet.Value,
					Amount:        0,
					Status:        "",
					StatusCode:    statusCode,
					StatusDetail:  statusText,
					Subject:       smsIsActive,
					IpAddress:     "",
					Payload:       util.TrimByteToString(isActiveMT),
				},
			)

			/**
			 * IF UNREG
			 */
		} else if existSub.RowsAffected == 1 && util.FilterUnreg(req.Message) {
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
			submitedId := resXML.Body.SubmitedID
			statusCode := resXML.Body.Code
			statusText := resXML.Body.Text

			// Update subscriptions
			subHasActive.LatestStatus = "SUCCESS"
			subHasActive.LatestSubject = smsUnsub
			subHasActive.UnsubAt = time.Now()
			subHasActive.PurgeAt = time.Now()
			subHasActive.RenewalAt = time.Time{}
			subHasActive.RetryAt = time.Time{}
			subHasActive.IsPurge = false
			subHasActive.IsRetry = false
			subHasActive.IsActive = false
			database.Datasource.DB().Save(&subHasActive)

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
					StatusCode:    statusCode,
					StatusDetail:  statusText,
					Subject:       smsUnsub,
					IpAddress:     "",
					Payload:       util.TrimByteToString(unsubMT),
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

			/**
			 * IF REG & REG KEREN
			 */
		} else if (existSub.RowsAffected == 0 && nonActiveSub.RowsAffected == 1) && util.FilterReg(req.Message) {
			subInActive.Keyword = strings.ToUpper(req.Message)
			subInActive.Adnet = adnet.Value
			subInActive.IpAddress = req.IpAddress
			database.Datasource.DB().Save(&subInActive)

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
			submitedId := resXML.Body.SubmitedID
			statusCode := resXML.Body.Code
			statusText := resXML.Body.Text

			/**
			 * if success status code = 0
			 */
			if statusCode == 0 {
				// update subscription
				subInActive.LatestSubject = smsFirstpush
				subInActive.LatestStatus = "SUCCESS"
				subInActive.Adnet = adnet.Value
				subInActive.Amount = service.Charge
				subInActive.RenewalAt = time.Now().AddDate(0, 0, service.Day)
				subInActive.ChargeAt = time.Now()
				subInActive.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
				subInActive.Success = subInActive.Success + 1
				subInActive.IpAddress = ""
				subInActive.IsRetry = false
				subInActive.IsPurge = false
				subInActive.IsActive = true
				database.Datasource.DB().Save(&subInActive)

				// insert transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedId,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        service.Charge,
						Status:        "SUCCESS",
						StatusCode:    statusCode,
						StatusDetail:  statusText,
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
				submitedIdwelcome := res1XML.Body.SubmitedID
				statusCodewelcome := res1XML.Body.Code
				statusTextwelcome := res1XML.Body.Text

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedIdwelcome,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "",
						StatusCode:    statusCodewelcome,
						StatusDetail:  statusTextwelcome,
						Subject:       smsWelcome,
						IpAddress:     "",
						Payload:       util.TrimByteToString(welcomeMT),
					},
				)

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

			} else if statusCode == 52 {
				subInActive.LatestSubject = smsFirstpush
				subInActive.LatestStatus = "FAILED"
				subInActive.Adnet = adnet.Value
				subInActive.Amount = 0
				subInActive.RenewalAt = time.Now().AddDate(0, 0, 1)
				subInActive.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
				subInActive.IpAddress = ""
				subInActive.IsRetry = true
				subInActive.IsPurge = false
				subInActive.IsActive = true
				database.Datasource.DB().Save(&subInActive)

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedId,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "FAILED",
						StatusCode:    statusCode,
						StatusDetail:  statusText,
						Subject:       smsFirstpush,
						IpAddress:     "",
						Payload:       util.TrimByteToString(firstpushMt),
					},
				)

				// sent mt_insuff
				insuffMT, err := handler.MessageTerminated(service, contInsuff, req.MobileNo, transactionId)
				if err != nil {
					loggerMt.WithFields(logrus.Fields{
						"transaction_id": transactionId,
						"msisdn":         req.MobileNo,
						"error":          err.Error(),
					}).Error(smsInsuff)
				}
				loggerMt.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"payload":        util.TrimByteToString(insuffMT),
				}).Info(smsInsuff)

				resultInsuff := util.EscapeChar(insuffMT)
				res1XML := dto.Response{}
				xml.Unmarshal([]byte(resultInsuff), &res1XML)
				submitedIdInsuff := resXML.Body.SubmitedID
				statusCodeInsuft := resXML.Body.Code
				statusTextInsuff := resXML.Body.Text

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedIdInsuff,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "",
						StatusCode:    statusCodeInsuft,
						StatusDetail:  statusTextInsuff,
						Subject:       smsInsuff,
						IpAddress:     "",
						Payload:       util.TrimByteToString(insuffMT),
					},
				)
			} else {
				subInActive.LatestSubject = smsFirstpush
				subInActive.LatestStatus = "FAILED"
				subInActive.Adnet = adnet.Value
				subInActive.IpAddress = ""
				subInActive.IsRetry = false
				subInActive.IsPurge = false
				subInActive.IsActive = false
				database.Datasource.DB().Save(&subInActive)

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedId,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "FAILED",
						StatusCode:    statusCode,
						StatusDetail:  statusText,
						Subject:       smsFirstpush,
						IpAddress:     "",
						Payload:       util.TrimByteToString(firstpushMt),
					},
				)
			}

			/**
			 * Postback
			 */
			postback, err := handler.Postback(service, req.MobileNo, adnet.Value, transactionId)
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
			 * IF UNREG
			 */
		} else if (existSub.RowsAffected == 0 && nonActiveSub.RowsAffected == 1) && util.FilterUnreg(req.Message) {

			// sent mt_purge
			purgeMT, err := handler.MessageTerminated(service, contPurge, req.MobileNo, transactionId)
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
				"payload":        util.TrimByteToString(purgeMT),
			}).Info(smsUnsub)

			resultPurge := util.EscapeChar(purgeMT)
			resXML := dto.Response{}
			xml.Unmarshal([]byte(resultPurge), &resXML)
			submitedId := resXML.Body.SubmitedID
			statusCode := resXML.Body.Code
			statusText := resXML.Body.Text

			// Insert to Transaction
			database.Datasource.DB().Create(
				&model.Transaction{
					TransactionID: transactionId,
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					SubmitedID:    submitedId,
					Keyword:       strings.ToUpper(req.Message),
					Adnet:         adnet.Value,
					Amount:        0,
					Status:        "",
					StatusCode:    statusCode,
					StatusDetail:  statusText,
					Subject:       smsUnsub,
					IpAddress:     "",
					Payload:       util.TrimByteToString(purgeMT),
				},
			)
			/**
			 * REG & NEW INPUT MSISDN
			 */
		} else if (existSub.RowsAffected == 0 || nonActiveSub.RowsAffected == 0) && util.FilterReg(req.Message) {
			database.Datasource.DB().Create(
				&model.Subscription{
					ServiceID:     service.ID,
					Msisdn:        req.MobileNo,
					Keyword:       strings.ToUpper(req.Message),
					Adnet:         adnet.Value,
					LatestSubject: "INPUT_MSISDN",
					Amount:        0,
					IpAddress:     req.IpAddress,
					IsActive:      true,
				},
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
			submitedId := resXML.Body.SubmitedID
			statusCode := resXML.Body.Code
			statusText := resXML.Body.Text

			var subscription model.Subscription
			database.Datasource.DB().
				Where("service_id", service.ID).
				Where("msisdn", req.MobileNo).
				Where("latest_subject", "INPUT_MSISDN").
				Where("is_active", true).
				First(&subscription)

			/**
			 * if success status code = 0
			 */
			if statusCode == 0 {

				// update subscription
				subscription.LatestSubject = smsFirstpush
				subscription.LatestStatus = "SUCCESS"
				subscription.Adnet = adnet.Value
				subscription.Amount = service.Charge
				subscription.RenewalAt = time.Now().AddDate(0, 0, service.Day)
				subscription.ChargeAt = time.Now()
				subscription.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
				subscription.Success = subscription.Success + 1
				subscription.IpAddress = ""
				subscription.IsRetry = false
				subscription.IsPurge = false
				subscription.IsActive = true
				database.Datasource.DB().Save(&subscription)

				// insert transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedId,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        service.Charge,
						Status:        "SUCCESS",
						StatusCode:    statusCode,
						StatusDetail:  statusText,
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
				submitedIdwelcome := res1XML.Body.SubmitedID
				statusCodewelcome := res1XML.Body.Code
				statusTextwelcome := res1XML.Body.Text

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedIdwelcome,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "",
						StatusCode:    statusCodewelcome,
						StatusDetail:  statusTextwelcome,
						Subject:       smsWelcome,
						IpAddress:     "",
						Payload:       util.TrimByteToString(welcomeMT),
					},
				)

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

			} else if statusCode == 52 {
				subscription.LatestSubject = smsFirstpush
				subscription.LatestStatus = "FAILED"
				subscription.Adnet = adnet.Value
				subscription.Amount = 0
				subscription.RenewalAt = time.Now().AddDate(0, 0, 1)
				subInActive.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
				subscription.IpAddress = ""
				subscription.IsRetry = true
				subscription.IsPurge = false
				subscription.IsActive = true
				database.Datasource.DB().Save(&subscription)

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedId,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "FAILED",
						StatusCode:    statusCode,
						StatusDetail:  statusText,
						Subject:       smsFirstpush,
						IpAddress:     "",
						Payload:       util.TrimByteToString(firstpushMt),
					},
				)

				// sent mt_insuff
				insuffMT, err := handler.MessageTerminated(service, contInsuff, req.MobileNo, transactionId)
				if err != nil {
					loggerMt.WithFields(logrus.Fields{
						"transaction_id": transactionId,
						"msisdn":         req.MobileNo,
						"error":          err.Error(),
					}).Error(smsInsuff)
				}
				loggerMt.WithFields(logrus.Fields{
					"transaction_id": transactionId,
					"msisdn":         req.MobileNo,
					"payload":        util.TrimByteToString(insuffMT),
				}).Info(smsInsuff)

				resultInsuff := util.EscapeChar(insuffMT)
				res1XML := dto.Response{}
				xml.Unmarshal([]byte(resultInsuff), &res1XML)
				submitedIdInsuff := resXML.Body.SubmitedID
				statusCodeInsuft := resXML.Body.Code
				statusTextInsuff := resXML.Body.Text

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedIdInsuff,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "",
						StatusCode:    statusCodeInsuft,
						StatusDetail:  statusTextInsuff,
						Subject:       smsInsuff,
						IpAddress:     "",
						Payload:       util.TrimByteToString(insuffMT),
					},
				)
			} else {
				subscription.LatestSubject = smsFirstpush
				subscription.LatestStatus = "FAILED"
				subscription.Adnet = adnet.Value
				subscription.Amount = 0
				subscription.RenewalAt = time.Time{}
				subscription.PurgeAt = time.Time{}
				subscription.IpAddress = ""
				subscription.IsRetry = false
				subscription.IsPurge = false
				subscription.IsActive = false
				database.Datasource.DB().Save(&subscription)

				// Insert to Transaction
				database.Datasource.DB().Create(
					&model.Transaction{
						TransactionID: transactionId,
						ServiceID:     service.ID,
						Msisdn:        req.MobileNo,
						SubmitedID:    submitedId,
						Keyword:       strings.ToUpper(req.Message),
						Adnet:         adnet.Value,
						Amount:        0,
						Status:        "FAILED",
						StatusCode:    statusCode,
						StatusDetail:  statusText,
						Subject:       smsFirstpush,
						IpAddress:     "",
						Payload:       util.TrimByteToString(firstpushMt),
					},
				)
			}

			/**
			 * Postback
			 */
			postback, err := handler.Postback(service, req.MobileNo, adnet.Value, transactionId)
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

		} else if (existSub.RowsAffected == 0 || nonActiveSub.RowsAffected == 0) && util.FilterUnreg(req.Message) {

			// sent mt_purge
			purgeMT, err := handler.MessageTerminated(service, contPurge, req.MobileNo, transactionId)
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
				"payload":        util.TrimByteToString(purgeMT),
			}).Info(smsUnsub)

			resultPurge := util.EscapeChar(purgeMT)
			resXML := dto.Response{}
			xml.Unmarshal([]byte(resultPurge), &resXML)
			submitedId := resXML.Body.SubmitedID
			statusCode := resXML.Body.Code
			statusText := resXML.Body.Text

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
					StatusCode:    statusCode,
					StatusDetail:  statusText,
					Subject:       smsUnsub,
					IpAddress:     "",
					Payload:       util.TrimByteToString(purgeMT),
				},
			)

		} else {
			/**
			 * IF WRONGKEY
			 */

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

			resultWrongkey := util.EscapeChar(wrongKeywordMt)
			resXML := dto.Response{}
			xml.Unmarshal([]byte(resultWrongkey), &resXML)
			submitedId := resXML.Body.SubmitedID
			statusCode := resXML.Body.Code
			statusText := resXML.Body.Text

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
					StatusCode:    statusCode,
					StatusDetail:  statusText,
					Subject:       smsWrongKey,
					IpAddress:     "",
					Payload:       util.TrimByteToString(wrongKeywordMt),
				},
			)
		}
	}

	wg.Done()
}

func drProccesor(wg *sync.WaitGroup, message []byte) {
	/**
	 * Sample Request
	 * {"msisdn":"62895635121559","shortcode":"99879","status":"DELIVRD","message":"1601666764269215859","ip":"116.206.10.222"}
	 */

	// parsing string json
	var req dto.DRRequest
	json.Unmarshal(message, &req)

	var transaction model.Transaction
	existTrans := database.Datasource.DB().Where("msisdn", req.Msisdn).Where("submited_id", req.Message).First(&transaction)

	if existTrans.RowsAffected == 1 {

		var labelStatus string
		if req.Status == "DELIVRD" {
			labelStatus = "SUCCESS"
		} else {
			labelStatus = "FAILED"
		}

		transaction.Status = labelStatus
		transaction.DrStatus = req.Status
		transaction.DrStatusDetail = util.DRStatus(req.Status)
		database.Datasource.DB().Save(&transaction)
	}

	wg.Done()
}

func renewalProccesor(wg *sync.WaitGroup, message []byte) {

	loggerMt := util.MakeLogger("mt", true)
	loggerNotif := util.MakeLogger("notif", true)

	transactionId := util.GenerateTransactionId()

	// parsing string json
	var sub model.Subscription
	json.Unmarshal(message, &sub)

	// get service by id
	service, _ := query.GetServiceById(sub.ServiceID)

	/**
	 * Query Content wording
	 */
	contRenewal, _ := query.GetContent(sub.ServiceID, "RENEWAL")
	// replaceRenewal := strings.NewReplacer("@purge_date", sub.PurgeAt.Format("02-Jan-2006"))
	// messageRenewal := replaceRenewal.Replace(contRenewal.Value)

	// sent mt_renewal
	renewalMt, err := handler.MessageTerminatedRenewal(service, contRenewal, sub.Msisdn, transactionId)
	if err != nil {
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         sub.Msisdn,
			"error":          err.Error(),
		}).Error(smsRenewal)
	}
	loggerMt.WithFields(logrus.Fields{
		"transaction_id": transactionId,
		"msisdn":         sub.Msisdn,
		"payload":        util.TrimByteToString(renewalMt),
	}).Info(smsRenewal)

	var (
		submitedId = ""
		statusCode = 0
		statusText = ""
	)

	if !json.Valid(renewalMt) {
		resultRenewal := util.EscapeChar(renewalMt)
		resXML := dto.Response{}
		xml.Unmarshal([]byte(resultRenewal), &resXML)
		submitedId = resXML.Body.SubmitedID
		statusCode = resXML.Body.Code
		statusText = resXML.Body.Text
	} else {
		resJSON := dto.ResponseJSON{}
		json.Unmarshal(renewalMt, &resJSON)
		submitedId = resJSON.Responses.ResponseBody.SubmitedID
		statusCode = resJSON.Responses.ResponseBody.Code
		statusText = resJSON.Responses.ResponseBody.Text
	}

	/**
	 * if success statusText = Successful
	 */
	if statusText == "Successful" {

		// Insert
		query.InsertTransact(database.Datasource.SqlDB(),
			model.Transaction{
				TransactionID: transactionId,
				ServiceID:     sub.ServiceID,
				Msisdn:        sub.Msisdn,
				SubmitedID:    submitedId,
				Keyword:       sub.Keyword,
				Subject:       smsRenewal,
				Amount:        service.Charge,
				Status:        "SUCCESS",
				StatusCode:    statusCode,
				StatusDetail:  statusText,
				IpAddress:     sub.IpAddress,
				Payload:       util.TrimByteToString(renewalMt),
			},
		)

		// Update last_subject, amount, renewal_at, charge_at, success, is_retry on subscription
		query.SubUpdateSuccess(database.Datasource.SqlDB(),
			model.Subscription{
				LatestSubject: smsRenewal,
				LatestStatus:  "SUCCESS",
				Amount:        service.Charge,
				RenewalAt:     time.Now().AddDate(0, 0, service.Day),
				ChargeAt:      time.Now(),
				Success:       1,
				IsRetry:       false,
				ServiceID:     sub.ServiceID,
				Msisdn:        sub.Msisdn,
			},
		)

		/**
		 * Notif Renewal
		 */
		notifRenewal, err := handler.NotifRenewal(service, sub.Msisdn, transactionId)
		if err != nil {
			loggerNotif.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         sub.Msisdn,
				"error":          err.Error(),
			}).Error()
		}
		loggerNotif.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         sub.Msisdn,
			"payload":        util.TrimByteToString(notifRenewal),
		}).Info()

	} else {

		query.InsertTransact(database.Datasource.SqlDB(),
			model.Transaction{
				TransactionID: transactionId,
				ServiceID:     sub.ServiceID,
				Msisdn:        sub.Msisdn,
				SubmitedID:    submitedId,
				Keyword:       sub.Keyword,
				Subject:       smsRenewal,
				Amount:        0,
				Status:        "FAILED",
				StatusCode:    statusCode,
				StatusDetail:  statusText,
				IpAddress:     sub.IpAddress,
				Payload:       util.TrimByteToString(renewalMt),
			},
		)

		// Update last_subject, amount, retry_at, is_retry on subscription
		query.SubUpdateFailed(database.Datasource.SqlDB(),
			model.Subscription{
				LatestSubject: smsRenewal,
				LatestStatus:  "FAILED",
				RenewalAt:     time.Now().AddDate(0, 0, 1),
				IsRetry:       true,
				ServiceID:     sub.ServiceID,
				Msisdn:        sub.Msisdn,
			},
		)
	}

	wg.Done()
}

func retryProccesor(wg *sync.WaitGroup, message []byte) {
	loggerMt := util.MakeLogger("mt", true)
	loggerNotif := util.MakeLogger("notif", true)

	transactionId := util.GenerateTransactionId()

	// parsing string json
	var sub model.Subscription
	json.Unmarshal(message, &sub)

	// get service by id
	service, _ := query.GetServiceById(sub.ServiceID)

	/**
	 * Query Content wording
	 */
	contRenewal, _ := query.GetContent(sub.ServiceID, "RENEWAL")
	// replaceRenewal := strings.NewReplacer("@purge_date", sub.PurgeAt.Format("02-Jan-2006"))
	// messageRenewal := replaceRenewal.Replace(contRenewal.Value)

	retryMt, err := handler.MessageTerminatedRenewal(service, contRenewal, sub.Msisdn, transactionId)
	if err != nil {
		loggerMt.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         sub.Msisdn,
			"error":          err.Error(),
		}).Error(smsFirstpush)
	}
	loggerMt.WithFields(logrus.Fields{
		"transaction_id": transactionId,
		"msisdn":         sub.Msisdn,
		"payload":        util.TrimByteToString(retryMt),
	}).Info(smsFirstpush)

	var (
		submitedId = ""
		statusCode = 0
		statusText = ""
	)

	if !json.Valid(retryMt) {
		resultRetry := util.EscapeChar(retryMt)
		resXML := dto.Response{}
		xml.Unmarshal([]byte(resultRetry), &resXML)
		submitedId = resXML.Body.SubmitedID
		statusCode = resXML.Body.Code
		statusText = resXML.Body.Text
	} else {
		resJSON := dto.ResponseJSON{}
		json.Unmarshal(retryMt, &resJSON)
		submitedId = resJSON.Responses.ResponseBody.SubmitedID
		statusCode = resJSON.Responses.ResponseBody.Code
		statusText = resJSON.Responses.ResponseBody.Text
	}

	/**
	 * if success statusText = Successful
	 */
	if statusText == "Successful" {
		query.RemoveTransact(database.Datasource.SqlDB(),
			model.Transaction{
				ServiceID: sub.ServiceID,
				Msisdn:    sub.Msisdn,
				Subject:   smsRenewal,
				Status:    "SUCCESS",
			},
		)

		// Insert new record if charging renewal success
		query.InsertTransact(database.Datasource.SqlDB(),
			model.Transaction{
				TransactionID: transactionId,
				ServiceID:     sub.ServiceID,
				Msisdn:        sub.Msisdn,
				SubmitedID:    submitedId,
				Keyword:       sub.Keyword,
				Subject:       smsRenewal,
				Amount:        service.Charge,
				Status:        "SUCCESS",
				StatusCode:    statusCode,
				StatusDetail:  statusText,
				IpAddress:     sub.IpAddress,
				Payload:       util.TrimByteToString(retryMt),
			},
		)

		// Update last_subject, amount, renewal_at, charge_at, success, is_retry on subscription
		query.SubUpdateSuccess(database.Datasource.SqlDB(),
			model.Subscription{
				LatestSubject: smsRenewal,
				LatestStatus:  "SUCCESS",
				Amount:        service.Charge,
				RenewalAt:     time.Now().AddDate(0, 0, service.Day),
				ChargeAt:      time.Now(),
				Success:       1,
				IsRetry:       false,
				ServiceID:     sub.ServiceID,
				Msisdn:        sub.Msisdn,
			},
		)

		/**
		 * Notif Renewal
		 */
		notifRenewal, err := handler.NotifRenewal(service, sub.Msisdn, transactionId)
		if err != nil {
			loggerNotif.WithFields(logrus.Fields{
				"transaction_id": transactionId,
				"msisdn":         sub.Msisdn,
				"error":          err.Error(),
			}).Error()
		}
		loggerNotif.WithFields(logrus.Fields{
			"transaction_id": transactionId,
			"msisdn":         sub.Msisdn,
			"payload":        util.TrimByteToString(notifRenewal),
		}).Info()
	}

	wg.Done()
}

func purgeProccesor(wg *sync.WaitGroup, message []byte) {

	// // parsing string json
	// var sub model.Subscription
	// json.Unmarshal(message, &sub)

	// // get service by id
	// service, _ := query.GetServiceById(sub.ServiceID)

	// var subscription model.Subscription
	// existSub := database.Datasource.DB().Where("service_id", service.ID).Where("msisdn", sub.Msisdn).First(&subscription)

	// if existSub.RowsAffected == 1 {
	// 	subscription.LatestSubject = smsPurge
	// 	subscription.LatestStatus = "SUCCESS"
	// 	subscription.IsPurge = true
	// 	subscription.IsActive = false
	// 	database.Datasource.DB().Save(&subscription)
	// }

	// wg.Done()
}
