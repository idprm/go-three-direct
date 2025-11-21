package handler

import (
	"encoding/json"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/pkg/util"
	"github.com/idprm/go-three-direct/internal/providers/telco"
	"github.com/idprm/go-three-direct/internal/services"
)

type RetryHandler struct {
	logger              *logger.Logger
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	sub                 *entity.Subscription
}

func NewRetryHandler(
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	sub *entity.Subscription,
) *RetryHandler {
	return &RetryHandler{
		logger:              logger,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		sub:                 sub,
	}
}

func (h *RetryHandler) Firstpush() error {

	trxId := util.GenerateTransactionId()

	service, err := h.serviceService.GetById(h.sub.ServiceID)
	if err != nil {
		return err
	}
	content, err := h.getFirstpushContent(h.sub.ServiceID)
	if err != nil {
		return err
	}

	p := telco.NewTelco(h.logger, service, content, h.sub)

	mt, err := p.MobileTerminated()
	if err != nil {
		return err
	}

	var (
		submitedId = ""
		statusCode = 0
		statusText = ""
	)

	if !json.Valid(mt) {

	} else {

	}

	/**
	 * if success statusText = Successful
	 */
	if statusCode == 0 && statusText == "Successful" {
		s := &entity.Subscription{}
		h.subscriptionService.Update(s)

		t := &entity.Transaction{
			TransactionID: trxId,
			ServiceID:     h.sub.ServiceID,
			Msisdn:        h.sub.Msisdn,
			SubmitedID:    submitedId,
			Keyword:       h.sub.Keyword,
			Subject:       SUBJECT_RENEWAL,
			Amount:        service.GetCharge(),
			Status:        STATUS_SUCCESS,
			StatusCode:    statusCode,
			StatusDetail:  statusText,
			IpAddress:     h.sub.IpAddress,
			Payload:       util.TrimByteToString(mt),
		}
		h.transactionService.Save(t)

		return nil
	}

	return nil
}

func (h *RetryHandler) Dailypush() error {
	trxId := util.GenerateTransactionId()

	service, err := h.serviceService.GetById(h.sub.ServiceID)
	if err != nil {
		return err
	}
	content, err := h.getRenewalContent(h.sub.ServiceID)
	if err != nil {
		return err
	}

	p := telco.NewTelco(h.logger, service, content, h.sub)

	mt, err := p.MobileTerminated()
	if err != nil {
		return err
	}

	var (
		submitedId = ""
		statusCode = 0
		statusText = ""
	)

	if !json.Valid(mt) {

	} else {

	}

	/**
	 * if success statusText = Successful
	 */
	if statusCode == 0 && statusText == "Successful" {
		s := &entity.Subscription{}
		h.subscriptionService.Update(s)

		t := &entity.Transaction{
			TransactionID: trxId,
			ServiceID:     h.sub.ServiceID,
			Msisdn:        h.sub.Msisdn,
			SubmitedID:    submitedId,
			Keyword:       h.sub.Keyword,
			Subject:       SUBJECT_RENEWAL,
			Amount:        service.GetCharge(),
			Status:        STATUS_SUCCESS,
			StatusCode:    statusCode,
			StatusDetail:  statusText,
			IpAddress:     h.sub.IpAddress,
			Payload:       util.TrimByteToString(mt),
		}
		h.transactionService.Save(t)

		return nil
	}

	return nil
}

func (h *RetryHandler) getFirstpushContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_FIRSTPUSH) {
		return nil, nil
	}
	return h.contentService.Get(serviceId, SMS_FIRSTPUSH)
}

func (h *RetryHandler) getRenewalContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_RENEWAL) {
		return nil, nil
	}
	return h.contentService.Get(serviceId, SMS_RENEWAL)
}

// loggerMt := util.MakeLogger("mt", true)
// 	loggerNotif := util.MakeLogger("notif", true)

// 	contentRepo := query.NewContentRepository(p.db)
// 	serviceRepo := query.NewServiceRepository(p.db)
// 	subscriptionRepo := query.NewSubscriptionRepository(p.db)
// 	transactionRepo := query.NewTransactionRepository(p.db)

// 	transactionId := util.GenerateTransactionId()

// 	// parsing string json
// 	var sub entity.Subscription
// 	json.Unmarshal(message, &sub)

// 	// get service by id
// 	service, _ := serviceRepo.GetServiceById(sub.ServiceID)

// 	var labelStatus string
// 	var contentStatus string
// 	if sub.IsCreatedAtToday() {
// 		labelStatus = smsFirstpush
// 		contentStatus = "FIRSTPUSH"
// 	} else {
// 		labelStatus = smsRenewal
// 		contentStatus = "RENEWAL"
// 	}

// 	/**
// 	 * Query Content wording
// 	 */
// 	content, _ := contentRepo.GetContent(sub.ServiceID, contentStatus)
// 	provider := handler.NewTelco()
// 	retryMt, err := provider.MessageTerminatedRenewal(service, content, sub.Msisdn, transactionId)
// 	if err != nil {
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         sub.Msisdn,
// 			"error":          err.Error(),
// 		}).Error(labelStatus)
// 	}
// 	loggerMt.WithFields(logrus.Fields{
// 		"transaction_id": transactionId,
// 		"msisdn":         sub.Msisdn,
// 		"payload":        util.TrimByteToString(retryMt),
// 	}).Info(labelStatus)

// 	var (
// 		submitedId = ""
// 		statusCode = 0
// 		statusText = ""
// 	)

// 	if !json.Valid(retryMt) {
// 		resultRetry := util.EscapeChar(retryMt)
// 		resXML := dto.Response{}
// 		xml.Unmarshal([]byte(resultRetry), &resXML)
// 		submitedId = resXML.Body.SubmitedID
// 		statusCode = resXML.Body.Code
// 		statusText = resXML.Body.Text
// 	} else {
// 		resJSON := dto.ResponseJSON{}
// 		json.Unmarshal(retryMt, &resJSON)
// 		submitedId = resJSON.Responses.ResponseBody.SubmitedID
// 		statusCode = resJSON.Responses.ResponseBody.Code
// 		statusText = resJSON.Responses.ResponseBody.Text
// 	}

// 	/**
// 	 * if success statusText = Successful
// 	 */
// 	if statusCode == 0 && statusText == "Successful" {
// 		transactionRepo.RemoveTransact(
// 			entity.Transaction{
// 				ServiceID: sub.ServiceID,
// 				Msisdn:    sub.Msisdn,
// 				Subject:   labelStatus,
// 				Status:    "SUCCESS",
// 			},
// 		)

// 		// Insert new record if charging renewal success
// 		transactionRepo.InsertTransact(
// 			entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     sub.ServiceID,
// 				Msisdn:        sub.Msisdn,
// 				SubmitedID:    submitedId,
// 				Keyword:       sub.Keyword,
// 				Subject:       labelStatus,
// 				Amount:        service.Charge,
// 				Status:        "SUCCESS",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				IpAddress:     sub.IpAddress,
// 				Payload:       util.TrimByteToString(retryMt),
// 			},
// 		)

// 		// Update last_subject, amount, renewal_at, charge_at, success, is_retry on subscription
// 		subscriptionRepo.SubUpdateSuccess(
// 			entity.Subscription{
// 				LatestSubject: labelStatus,
// 				LatestStatus:  "SUCCESS",
// 				Amount:        service.Charge,
// 				RenewalAt:     time.Now().AddDate(0, 0, service.Day),
// 				ChargeAt:      time.Now(),
// 				Success:       1,
// 				IsRetry:       false,
// 				ServiceID:     sub.ServiceID,
// 				Msisdn:        sub.Msisdn,
// 			},
// 		)

// 		/**
// 		 * Notif Renewal
// 		 */
// 		notifRenewal, err := handler.NotifRenewal(service, sub.Msisdn, transactionId)
// 		if err != nil {
// 			loggerNotif.WithFields(logrus.Fields{
// 				"transaction_id": transactionId,
// 				"msisdn":         sub.Msisdn,
// 				"error":          err.Error(),
// 			}).Error()
// 		}
// 		loggerNotif.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         sub.Msisdn,
// 			"payload":        util.TrimByteToString(notifRenewal),
// 		}).Info()
// 	}
