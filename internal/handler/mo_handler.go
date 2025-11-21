package handler

import (
	"fmt"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/model"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/services"
)

type MOHandler struct {
	logger              *logger.Logger
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	req                 *model.MORequest
}

func NewMOHandler(
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	req *model.MORequest,
) *MOHandler {
	return &MOHandler{
		logger:              logger,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		req:                 req,
	}
}

func (h *MOHandler) Register() error {

	return nil
}

func (h *MOHandler) Firstpush() error {

	return nil
}

func (h *MOHandler) IsActive() bool {
	return h.subscriptionService.IsActive(1, "")
}

func (h *MOHandler) IsReg() bool {
	return true
}

func (h *MOHandler) IsUnreg() bool {
	return true
}

func (h *MOHandler) getFirstpushContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_FIRSTPUSH) {
		return nil, fmt.Errorf("firstpush content not found for service ID %d", serviceId)
	}
	return h.contentService.Get(serviceId, SMS_FIRSTPUSH)
}

func (h *MOHandler) getWelcomeContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_WELCOME) {
		return nil, fmt.Errorf("welcome content not found for service ID %d", serviceId)
	}
	return h.contentService.Get(serviceId, SMS_WELCOME)
}

func (h *MOHandler) getInsuffContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_INSUFF) {
		return nil, fmt.Errorf("insuff content not found for service ID %d", serviceId)
	}
	return h.contentService.Get(serviceId, SMS_INSUFF)
}

func (h *MOHandler) getIsActiveContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_ISACTIVE) {
		return nil, fmt.Errorf("is_active content not found for service ID %d", serviceId)
	}
	return h.contentService.Get(serviceId, SMS_ISACTIVE)
}

func (h *MOHandler) getUnsubContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_UNSUB) {
		return nil, fmt.Errorf("unsub content not found for service ID %d", serviceId)
	}
	return h.contentService.Get(serviceId, SMS_UNSUB)
}

func (h *MOHandler) getPurgeContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_PURGE) {
		return nil, fmt.Errorf("purge content not found for service ID %d", serviceId)
	}
	return h.contentService.Get(serviceId, SMS_PURGE)
}

func (h *MOHandler) getWrongKeyContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_WRONGKEY) {
		return nil, fmt.Errorf("wrongkey content not found for service ID %d", serviceId)
	}
	return h.contentService.Get(serviceId, SMS_WRONGKEY)
}

// /**
// 	 * Sample Request
// 	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
// 	 */
// 	loggerMt := util.MakeLogger("mt", true)
// 	loggerNotif := util.MakeLogger("notif", true)
// 	loggerPb := util.MakeLogger("postback", true)

// 	transactionId := util.GenerateTransactionId()

// 	// parsing string json
// 	var req dto.MORequest
// 	json.Unmarshal(message, &req)

// 	/**
// 	 * Query content
// 	 */
// 	// get service by name
// 	service, err := serviceRepo.GetServiceByName(util.FilterMessage(strings.ToUpper(req.Message)))
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	provider := handler.NewTelco()

// 	/**
// 	 * Query Content
// 	 */
// 	contFirstpush, _ := contentRepo.GetContent(service.ID, valFirstpush)

// 	contWelcome, _ := contentRepo.GetContent(service.ID, valWelcome)

// 	contInsuff, _ := contentRepo.GetContent(service.ID, valInsuft)

// 	contIsActive, _ := contentRepo.GetContent(service.ID, valIsActive)

// 	contUnsub, _ := contentRepo.GetContent(service.ID, valUnsub)

// 	contPurge, _ := contentRepo.GetContent(service.ID, valPurge)

// 	contWrongKey, _ := contentRepo.GetContent(service.ID, valErroyKey)

// 	var subHasActive entity.Subscription
// 	existSub := p.gdb.Where("service_id", service.ID).Where("msisdn", req.MobileNo).Where("is_active", true).First(&subHasActive)

// 	var subInActive entity.Subscription
// 	nonActiveSub := p.gdb.Where("service_id", service.ID).Where("msisdn", req.MobileNo).Where("is_active", false).First(&subInActive)

// 	adn := util.KeywordDefine(strings.ToUpper(req.Message))

// 	var adnet entity.Adnet
// 	p.gdb.Where("name", adn).First(&adnet)

// 	/**
// 	 * IF SUB IS EXIST AND IS_ACTIVE = true
// 	 */
// 	if existSub.RowsAffected == 1 && util.FilterReg(req.Message) {
// 		subHasActive.Keyword = strings.ToUpper(req.Message)
// 		subHasActive.Adnet = adnet.Value
// 		subHasActive.IpAddress = req.IpAddress
// 		p.gdb.Save(&subHasActive)

// 		// sent mt_is_active
// 		isActiveMT, err := provider.MessageTerminated(service, contIsActive, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error(smsIsActive)
// 		}
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(isActiveMT),
// 		}).Info(smsIsActive)

// 		resultIsActive := util.EscapeChar(isActiveMT)
// 		resXML := dto.Response{}
// 		xml.Unmarshal([]byte(resultIsActive), &resXML)
// 		submitedId := resXML.Body.SubmitedID
// 		statusCode := resXML.Body.Code
// 		statusText := resXML.Body.Text

// 		// Insert to Transaction
// 		p.gdb.Create(
// 			&entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     service.ID,
// 				Msisdn:        req.MobileNo,
// 				SubmitedID:    submitedId,
// 				Keyword:       strings.ToUpper(req.Message),
// 				Adnet:         adnet.Value,
// 				Amount:        0,
// 				Status:        "",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				Subject:       smsIsActive,
// 				IpAddress:     "",
// 				Payload:       util.TrimByteToString(isActiveMT),
// 			},
// 		)

// 		/**
// 		 * IF UNREG
// 		 */
// 	} else if existSub.RowsAffected == 1 && util.FilterUnreg(req.Message) {
// 		// sent mt_unsub
// 		unsubMT, err := provider.MessageTerminated(service, contUnsub, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error(smsUnsub)
// 		}
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(unsubMT),
// 		}).Info(smsUnsub)

// 		resultUnsub := util.EscapeChar(unsubMT)
// 		resXML := dto.Response{}
// 		xml.Unmarshal([]byte(resultUnsub), &resXML)
// 		submitedId := resXML.Body.SubmitedID
// 		statusCode := resXML.Body.Code
// 		statusText := resXML.Body.Text

// 		// Update subscriptions
// 		subHasActive.LatestStatus = "SUCCESS"
// 		subHasActive.LatestSubject = smsUnsub
// 		subHasActive.UnsubAt = time.Now()
// 		subHasActive.PurgeAt = time.Now()
// 		subHasActive.RenewalAt = time.Time{}
// 		subHasActive.RetryAt = time.Time{}
// 		subHasActive.IsPurge = false
// 		subHasActive.IsRetry = false
// 		subHasActive.IsActive = false
// 		p.gdb.Save(&subHasActive)

// 		// Insert to Transaction
// 		p.gdb.Create(
// 			&entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     service.ID,
// 				Msisdn:        req.MobileNo,
// 				SubmitedID:    submitedId,
// 				Keyword:       strings.ToUpper(req.Message),
// 				Amount:        0,
// 				Status:        "",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				Subject:       smsUnsub,
// 				IpAddress:     "",
// 				Payload:       util.TrimByteToString(unsubMT),
// 			},
// 		)

// 		/**
// 		 * Notif Unsub
// 		 */
// 		notifUnsub, err := handler.NotifUnsub(service, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerNotif.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error()
// 		}
// 		loggerNotif.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(notifUnsub),
// 		}).Info()

// 		/**
// 		 * IF REG & REG KEREN
// 		 */
// 	} else if (existSub.RowsAffected == 0 && nonActiveSub.RowsAffected == 1) && util.FilterReg(req.Message) {
// 		subInActive.Keyword = strings.ToUpper(req.Message)
// 		subInActive.Adnet = adnet.Value
// 		subInActive.IpAddress = req.IpAddress
// 		p.gdb.Save(&subInActive)

// 		// sent mt_firstpush
// 		firstpushMt, err := provider.MessageTerminated(service, contFirstpush, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error(smsFirstpush)
// 		}
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(firstpushMt),
// 		}).Info(smsFirstpush)

// 		var (
// 			submitedId = ""
// 			statusCode = 0
// 			statusText = ""
// 		)

// 		if !json.Valid(firstpushMt) {
// 			resultFirstpush := util.EscapeChar(firstpushMt)
// 			resXML := dto.Response{}
// 			xml.Unmarshal([]byte(resultFirstpush), &resXML)
// 			submitedId = resXML.Body.SubmitedID
// 			statusCode = resXML.Body.Code
// 			statusText = resXML.Body.Text
// 		} else {
// 			resJSON := dto.ResponseJSON{}
// 			json.Unmarshal(firstpushMt, &resJSON)
// 			submitedId = resJSON.Responses.ResponseBody.SubmitedID
// 			statusCode = resJSON.Responses.ResponseBody.Code
// 			statusText = resJSON.Responses.ResponseBody.Text
// 		}

// 		/**
// 		 * if success status code = 0
// 		 */
// 		if statusCode == 0 && statusText == "Successful" {
// 			// update subscription
// 			subInActive.LatestSubject = smsFirstpush
// 			subInActive.LatestStatus = "SUCCESS"
// 			subInActive.Adnet = adnet.Value
// 			subInActive.Amount = service.Charge
// 			subInActive.RenewalAt = time.Now().AddDate(0, 0, service.Day)
// 			subInActive.ChargeAt = time.Now()
// 			subInActive.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
// 			subInActive.Success = subInActive.Success + 1
// 			subInActive.IpAddress = ""
// 			subInActive.IsRetry = false
// 			subInActive.IsPurge = false
// 			subInActive.IsActive = true
// 			p.gdb.Save(&subInActive)

// 			// insert transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedId,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        service.Charge,
// 					Status:        "SUCCESS",
// 					StatusCode:    statusCode,
// 					StatusDetail:  statusText,
// 					Subject:       smsFirstpush,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(firstpushMt),
// 				},
// 			)

// 			// sent mt_welcome
// 			welcomeMT, err := provider.MessageTerminated(service, contWelcome, req.MobileNo, transactionId)
// 			if err != nil {
// 				loggerMt.WithFields(logrus.Fields{
// 					"transaction_id": transactionId,
// 					"msisdn":         req.MobileNo,
// 					"error":          err.Error(),
// 				}).Error(smsWelcome)
// 			}
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"payload":        util.TrimByteToString(welcomeMT),
// 			}).Info(smsWelcome)

// 			resultWelcome := util.EscapeChar(welcomeMT)
// 			res1XML := dto.Response{}
// 			xml.Unmarshal([]byte(resultWelcome), &res1XML)
// 			submitedIdwelcome := res1XML.Body.SubmitedID
// 			statusCodewelcome := res1XML.Body.Code
// 			statusTextwelcome := res1XML.Body.Text

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedIdwelcome,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "",
// 					StatusCode:    statusCodewelcome,
// 					StatusDetail:  statusTextwelcome,
// 					Subject:       smsWelcome,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(welcomeMT),
// 				},
// 			)

// 			/**
// 			 * Notif sub
// 			 */
// 			notifSub, err := handler.NotifSub(service, req.MobileNo, transactionId)
// 			if err != nil {
// 				loggerNotif.WithFields(logrus.Fields{
// 					"transaction_id": transactionId,
// 					"msisdn":         req.MobileNo,
// 					"error":          err.Error(),
// 				}).Error()
// 			}
// 			loggerNotif.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"payload":        util.TrimByteToString(notifSub),
// 			}).Info()

// 		} else if statusCode == 52 {
// 			subInActive.LatestSubject = smsFirstpush
// 			subInActive.LatestStatus = "FAILED"
// 			subInActive.Adnet = adnet.Value
// 			subInActive.Amount = 0
// 			subInActive.RenewalAt = time.Now().AddDate(0, 0, 1)
// 			subInActive.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
// 			subInActive.IpAddress = ""
// 			subInActive.IsRetry = true
// 			subInActive.IsPurge = false
// 			subInActive.IsActive = true
// 			p.gdb.Save(&subInActive)

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedId,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "FAILED",
// 					StatusCode:    statusCode,
// 					StatusDetail:  statusText,
// 					Subject:       smsFirstpush,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(firstpushMt),
// 				},
// 			)

// 			// sent mt_insuff
// 			insuffMT, err := provider.MessageTerminated(service, contInsuff, req.MobileNo, transactionId)
// 			if err != nil {
// 				loggerMt.WithFields(logrus.Fields{
// 					"transaction_id": transactionId,
// 					"msisdn":         req.MobileNo,
// 					"error":          err.Error(),
// 				}).Error(smsInsuff)
// 			}
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"payload":        util.TrimByteToString(insuffMT),
// 			}).Info(smsInsuff)

// 			resultInsuff := util.EscapeChar(insuffMT)
// 			res1XML := dto.Response{}
// 			xml.Unmarshal([]byte(resultInsuff), &res1XML)
// 			submitedIdInsuff := res1XML.Body.SubmitedID
// 			statusCodeInsuft := res1XML.Body.Code
// 			statusTextInsuff := res1XML.Body.Text

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedIdInsuff,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "",
// 					StatusCode:    statusCodeInsuft,
// 					StatusDetail:  statusTextInsuff,
// 					Subject:       smsInsuff,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(insuffMT),
// 				},
// 			)
// 		} else {
// 			subInActive.LatestSubject = smsFirstpush
// 			subInActive.LatestStatus = "FAILED"
// 			subInActive.Adnet = adnet.Value
// 			subInActive.IpAddress = ""
// 			subInActive.IsRetry = false
// 			subInActive.IsPurge = false
// 			subInActive.IsActive = false
// 			p.gdb.Save(&subInActive)

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedId,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "FAILED",
// 					StatusCode:    statusCode,
// 					StatusDetail:  statusText,
// 					Subject:       smsFirstpush,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(firstpushMt),
// 				},
// 			)
// 		}

// 		/**
// 		 * Postback
// 		 */
// 		postback, err := handler.Postback(service, req.MobileNo, adnet.Value, transactionId)
// 		if err != nil {
// 			loggerPb.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error()
// 		}
// 		loggerPb.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(postback),
// 		}).Info()

// 		/**
// 		 * IF UNREG
// 		 */
// 	} else if (existSub.RowsAffected == 0 && nonActiveSub.RowsAffected == 1) && util.FilterUnreg(req.Message) {

// 		// sent mt_purge
// 		purgeMT, err := provider.MessageTerminated(service, contPurge, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error(smsUnsub)
// 		}
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(purgeMT),
// 		}).Info(smsUnsub)

// 		resultPurge := util.EscapeChar(purgeMT)
// 		resXML := dto.Response{}
// 		xml.Unmarshal([]byte(resultPurge), &resXML)
// 		submitedId := resXML.Body.SubmitedID
// 		statusCode := resXML.Body.Code
// 		statusText := resXML.Body.Text

// 		// Insert to Transaction
// 		p.gdb.Create(
// 			&entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     service.ID,
// 				Msisdn:        req.MobileNo,
// 				SubmitedID:    submitedId,
// 				Keyword:       strings.ToUpper(req.Message),
// 				Adnet:         adnet.Value,
// 				Amount:        0,
// 				Status:        "",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				Subject:       smsUnsub,
// 				IpAddress:     "",
// 				Payload:       util.TrimByteToString(purgeMT),
// 			},
// 		)
// 		/**
// 		 * REG & NEW INPUT MSISDN
// 		 */
// 	} else if (existSub.RowsAffected == 0 || nonActiveSub.RowsAffected == 0) && util.FilterReg(req.Message) {
// 		p.gdb.Create(
// 			&entity.Subscription{
// 				ServiceID:     service.ID,
// 				Msisdn:        req.MobileNo,
// 				Keyword:       strings.ToUpper(req.Message),
// 				Adnet:         adnet.Value,
// 				LatestSubject: "INPUT_MSISDN",
// 				Amount:        0,
// 				IpAddress:     req.IpAddress,
// 				IsActive:      true,
// 			},
// 		)

// 		// sent mt_firstpush
// 		firstpushMt, err := provider.MessageTerminated(service, contFirstpush, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error(smsFirstpush)
// 		}
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(firstpushMt),
// 		}).Info(smsFirstpush)

// 		var (
// 			submitedId = ""
// 			statusCode = 0
// 			statusText = ""
// 		)

// 		if !json.Valid(firstpushMt) {
// 			resultFirstpush := util.EscapeChar(firstpushMt)
// 			resXML := dto.Response{}
// 			xml.Unmarshal([]byte(resultFirstpush), &resXML)
// 			submitedId = resXML.Body.SubmitedID
// 			statusCode = resXML.Body.Code
// 			statusText = resXML.Body.Text
// 		} else {
// 			resJSON := dto.ResponseJSON{}
// 			json.Unmarshal(firstpushMt, &resJSON)
// 			submitedId = resJSON.Responses.ResponseBody.SubmitedID
// 			statusCode = resJSON.Responses.ResponseBody.Code
// 			statusText = resJSON.Responses.ResponseBody.Text
// 		}

// 		var subscription entity.Subscription
// 		p.gdb.
// 			Where("service_id", service.ID).
// 			Where("msisdn", req.MobileNo).
// 			Where("latest_subject", "INPUT_MSISDN").
// 			Where("is_active", true).
// 			First(&subscription)

// 		/**
// 		 * if success status code = 0
// 		 */
// 		if statusCode == 0 && statusText == "Successful" {
// 			// update subscription
// 			subscription.LatestSubject = smsFirstpush
// 			subscription.LatestStatus = "SUCCESS"
// 			subscription.Adnet = adnet.Value
// 			subscription.Amount = service.Charge
// 			subscription.RenewalAt = time.Now().AddDate(0, 0, service.Day)
// 			subscription.ChargeAt = time.Now()
// 			subscription.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
// 			subscription.Success = subscription.Success + 1
// 			subscription.IpAddress = ""
// 			subscription.IsRetry = false
// 			subscription.IsPurge = false
// 			subscription.IsActive = true
// 			p.gdb.Save(&subscription)

// 			// insert transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedId,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        service.Charge,
// 					Status:        "SUCCESS",
// 					StatusCode:    statusCode,
// 					StatusDetail:  statusText,
// 					Subject:       smsFirstpush,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(firstpushMt),
// 				},
// 			)

// 			// sent mt_welcome
// 			welcomeMT, err := provider.MessageTerminated(service, contWelcome, req.MobileNo, transactionId)
// 			if err != nil {
// 				loggerMt.WithFields(logrus.Fields{
// 					"transaction_id": transactionId,
// 					"msisdn":         req.MobileNo,
// 					"error":          err.Error(),
// 				}).Error(smsWelcome)
// 			}
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"payload":        util.TrimByteToString(welcomeMT),
// 			}).Info(smsWelcome)

// 			resultWelcome := util.EscapeChar(welcomeMT)
// 			res1XML := dto.Response{}
// 			xml.Unmarshal([]byte(resultWelcome), &res1XML)
// 			submitedIdwelcome := res1XML.Body.SubmitedID
// 			statusCodewelcome := res1XML.Body.Code
// 			statusTextwelcome := res1XML.Body.Text

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedIdwelcome,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "",
// 					StatusCode:    statusCodewelcome,
// 					StatusDetail:  statusTextwelcome,
// 					Subject:       smsWelcome,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(welcomeMT),
// 				},
// 			)

// 			/**
// 			 * Notif sub
// 			 */
// 			notifSub, err := handler.NotifSub(service, req.MobileNo, transactionId)
// 			if err != nil {
// 				loggerNotif.WithFields(logrus.Fields{
// 					"transaction_id": transactionId,
// 					"msisdn":         req.MobileNo,
// 					"error":          err.Error(),
// 				}).Error()
// 			}
// 			loggerNotif.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"payload":        util.TrimByteToString(notifSub),
// 			}).Info()

// 		} else if statusCode == 52 {
// 			subscription.LatestSubject = smsFirstpush
// 			subscription.LatestStatus = "FAILED"
// 			subscription.Adnet = adnet.Value
// 			subscription.Amount = 0
// 			subscription.RenewalAt = time.Now().AddDate(0, 0, 1)
// 			subInActive.PurgeAt = time.Now().AddDate(0, 0, service.PurgeDay)
// 			subscription.IpAddress = ""
// 			subscription.IsRetry = true
// 			subscription.IsPurge = false
// 			subscription.IsActive = true
// 			p.gdb.Save(&subscription)

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedId,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "FAILED",
// 					StatusCode:    statusCode,
// 					StatusDetail:  statusText,
// 					Subject:       smsFirstpush,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(firstpushMt),
// 				},
// 			)

// 			// sent mt_insuff
// 			insuffMT, err := provider.MessageTerminated(service, contInsuff, req.MobileNo, transactionId)
// 			if err != nil {
// 				loggerMt.WithFields(logrus.Fields{
// 					"transaction_id": transactionId,
// 					"msisdn":         req.MobileNo,
// 					"error":          err.Error(),
// 				}).Error(smsInsuff)
// 			}
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"payload":        util.TrimByteToString(insuffMT),
// 			}).Info(smsInsuff)

// 			resultInsuff := util.EscapeChar(insuffMT)
// 			res1XML := dto.Response{}
// 			xml.Unmarshal([]byte(resultInsuff), &res1XML)
// 			submitedIdInsuff := res1XML.Body.SubmitedID
// 			statusCodeInsuft := res1XML.Body.Code
// 			statusTextInsuff := res1XML.Body.Text

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedIdInsuff,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "",
// 					StatusCode:    statusCodeInsuft,
// 					StatusDetail:  statusTextInsuff,
// 					Subject:       smsInsuff,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(insuffMT),
// 				},
// 			)
// 		} else {
// 			subscription.LatestSubject = smsFirstpush
// 			subscription.LatestStatus = "FAILED"
// 			subscription.Adnet = adnet.Value
// 			subscription.Amount = 0
// 			subscription.RenewalAt = time.Time{}
// 			subscription.PurgeAt = time.Time{}
// 			subscription.IpAddress = ""
// 			subscription.IsRetry = false
// 			subscription.IsPurge = false
// 			subscription.IsActive = false
// 			p.gdb.Save(&subscription)

// 			// Insert to Transaction
// 			p.gdb.Create(
// 				&entity.Transaction{
// 					TransactionID: transactionId,
// 					ServiceID:     service.ID,
// 					Msisdn:        req.MobileNo,
// 					SubmitedID:    submitedId,
// 					Keyword:       strings.ToUpper(req.Message),
// 					Adnet:         adnet.Value,
// 					Amount:        0,
// 					Status:        "FAILED",
// 					StatusCode:    statusCode,
// 					StatusDetail:  statusText,
// 					Subject:       smsFirstpush,
// 					IpAddress:     "",
// 					Payload:       util.TrimByteToString(firstpushMt),
// 				},
// 			)
// 		}

// 		/**
// 		 * Postback
// 		 */
// 		postback, err := handler.Postback(service, req.MobileNo, adnet.Value, transactionId)
// 		if err != nil {
// 			loggerPb.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error()
// 		}
// 		loggerPb.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(postback),
// 		}).Info()

// 	} else if (existSub.RowsAffected == 0 || nonActiveSub.RowsAffected == 0) && util.FilterUnreg(req.Message) {

// 		// sent mt_purge
// 		purgeMT, err := provider.MessageTerminated(service, contPurge, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error(smsUnsub)
// 		}
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(purgeMT),
// 		}).Info(smsUnsub)

// 		resultPurge := util.EscapeChar(purgeMT)
// 		resXML := dto.Response{}
// 		xml.Unmarshal([]byte(resultPurge), &resXML)
// 		submitedId := resXML.Body.SubmitedID
// 		statusCode := resXML.Body.Code
// 		statusText := resXML.Body.Text

// 		// Insert to Transaction
// 		p.gdb.Create(
// 			&entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     service.ID,
// 				Msisdn:        req.MobileNo,
// 				SubmitedID:    submitedId,
// 				Keyword:       strings.ToUpper(req.Message),
// 				Amount:        0,
// 				Status:        "",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				Subject:       smsUnsub,
// 				IpAddress:     "",
// 				Payload:       util.TrimByteToString(purgeMT),
// 			},
// 		)

// 	} else {
// 		/**
// 		 * IF WRONGKEY
// 		 */

// 		// sent mt_wrongkey
// 		wrongKeywordMt, err := provider.MessageTerminated(service, contWrongKey, req.MobileNo, transactionId)
// 		if err != nil {
// 			loggerMt.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         req.MobileNo,
// 				"error":          err.Error(),
// 			}).Error(smsWrongKey)
// 		}
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         req.MobileNo,
// 			"payload":        util.TrimByteToString(wrongKeywordMt),
// 		}).Info(smsWrongKey)

// 		resultWrongkey := util.EscapeChar(wrongKeywordMt)
// 		resXML := dto.Response{}
// 		xml.Unmarshal([]byte(resultWrongkey), &resXML)
// 		submitedId := resXML.Body.SubmitedID
// 		statusCode := resXML.Body.Code
// 		statusText := resXML.Body.Text

// 		// Insert to Transaction
// 		p.gdb.Create(
// 			&entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     service.ID,
// 				Msisdn:        req.MobileNo,
// 				SubmitedID:    submitedId,
// 				Keyword:       strings.ToUpper(req.Message),
// 				Amount:        0,
// 				Status:        "",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				Subject:       smsWrongKey,
// 				IpAddress:     "",
// 				Payload:       util.TrimByteToString(wrongKeywordMt),
// 			},
// 		)
// 	}
