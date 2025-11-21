package handler

import (
	"encoding/json"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/pkg/util"
	"github.com/idprm/go-three-direct/internal/providers/telco"
	"github.com/idprm/go-three-direct/internal/services"
)

type RenewalHandler struct {
	logger              *logger.Logger
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	sub                 *entity.Subscription
}

func NewRenewalHandler(
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	sub *entity.Subscription,
) *RenewalHandler {
	return &RenewalHandler{
		logger:              logger,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		sub:                 sub,
	}
}

func (h *RenewalHandler) Dailypush() error {

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
	} else {
		t := &entity.Transaction{
			TransactionID: trxId,
			ServiceID:     h.sub.ServiceID,
			Msisdn:        h.sub.Msisdn,
			SubmitedID:    submitedId,
			Keyword:       h.sub.Keyword,
			Subject:       SUBJECT_RENEWAL,
			Amount:        0,
			Status:        STATUS_FAILED,
			StatusCode:    statusCode,
			StatusDetail:  statusText,
			IpAddress:     h.sub.IpAddress,
			Payload:       util.TrimByteToString(mt),
		}
		h.transactionService.Save(t)
	}

	// message terminated renewal logic here
	return nil
}

func (h *RenewalHandler) isRenewal() bool {
	return h.subscriptionService.IsRetry(h.sub.GetServiceId(), h.sub.GetMsisdn())
}

func (h *RenewalHandler) getRenewalContent(serviceId int) (*entity.Content, error) {
	if !h.contentService.IsContent(serviceId, SMS_RENEWAL) {
		return nil, nil
	}
	return h.contentService.Get(serviceId, SMS_RENEWAL)
}

// loggerMt := util.MakeLogger("mt", true)
// 	loggerNotif := util.MakeLogger("notif", true)

// 	transactionId := util.GenerateTransactionId()

// 	// parsing string json
// 	var sub entity.Subscription
// 	json.Unmarshal(message, &sub)

// 	// get service by id
// 	service, _ := serviceRepo.GetServiceById(sub.ServiceID)

// 	/**
// 	 * Query Content wording
// 	 */
// 	contRenewal, _ := contentRepo.GetContent(sub.ServiceID, "RENEWAL")
// 	// replaceRenewal := strings.NewReplacer("@purge_date", sub.PurgeAt.Format("02-Jan-2006"))
// 	// messageRenewal := replaceRenewal.Replace(contRenewal.Value)

// 	provider := handler.NewTelco()

// 	// sent mt_renewal
// 	renewalMt, err := provider.MessageTerminatedRenewal(service, contRenewal, sub.Msisdn, transactionId)
// 	if err != nil {
// 		loggerMt.WithFields(logrus.Fields{
// 			"transaction_id": transactionId,
// 			"msisdn":         sub.Msisdn,
// 			"error":          err.Error(),
// 		}).Error(smsRenewal)
// 	}
// 	loggerMt.WithFields(logrus.Fields{
// 		"transaction_id": transactionId,
// 		"msisdn":         sub.Msisdn,
// 		"payload":        util.TrimByteToString(renewalMt),
// 	}).Info(smsRenewal)

// 	var (
// 		submitedId = ""
// 		statusCode = 0
// 		statusText = ""
// 	)

// 	if !json.Valid(renewalMt) {
// 		resultRenewal := util.EscapeChar(renewalMt)
// 		resXML := dto.Response{}
// 		xml.Unmarshal([]byte(resultRenewal), &resXML)
// 		submitedId = resXML.Body.SubmitedID
// 		statusCode = resXML.Body.Code
// 		statusText = resXML.Body.Text
// 	} else {
// 		resJSON := dto.ResponseJSON{}
// 		json.Unmarshal(renewalMt, &resJSON)
// 		submitedId = resJSON.Responses.ResponseBody.SubmitedID
// 		statusCode = resJSON.Responses.ResponseBody.Code
// 		statusText = resJSON.Responses.ResponseBody.Text
// 	}

// 	/**
// 	 * if success statusText = Successful
// 	 */
// 	if statusCode == 0 && statusText == "Successful" {

// 		// Insert
// 		transactionRepo.InsertTransact(
// 			entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     sub.ServiceID,
// 				Msisdn:        sub.Msisdn,
// 				SubmitedID:    submitedId,
// 				Keyword:       sub.Keyword,
// 				Subject:       smsRenewal,
// 				Amount:        service.Charge,
// 				Status:        "SUCCESS",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				IpAddress:     sub.IpAddress,
// 				Payload:       util.TrimByteToString(renewalMt),
// 			},
// 		)

// 		// Update last_subject, amount, renewal_at, charge_at, success, is_retry on subscription
// 		subscriptionRepo.SubUpdateSuccess(
// 			entity.Subscription{
// 				LatestSubject: smsRenewal,
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

// 	} else {

// 		transactionRepo.InsertTransact(
// 			entity.Transaction{
// 				TransactionID: transactionId,
// 				ServiceID:     sub.ServiceID,
// 				Msisdn:        sub.Msisdn,
// 				SubmitedID:    submitedId,
// 				Keyword:       sub.Keyword,
// 				Subject:       smsRenewal,
// 				Amount:        0,
// 				Status:        "FAILED",
// 				StatusCode:    statusCode,
// 				StatusDetail:  statusText,
// 				IpAddress:     sub.IpAddress,
// 				Payload:       util.TrimByteToString(renewalMt),
// 			},
// 		)

// 		// Update last_subject, amount, retry_at, is_retry on subscription
// 		subscriptionRepo.SubUpdateFailed(
// 			entity.Subscription{
// 				LatestSubject: smsRenewal,
// 				LatestStatus:  "FAILED",
// 				RenewalAt:     time.Now().AddDate(0, 0, 1),
// 				IsRetry:       true,
// 				ServiceID:     sub.ServiceID,
// 				Msisdn:        sub.Msisdn,
// 			},
// 		)
// 	}
