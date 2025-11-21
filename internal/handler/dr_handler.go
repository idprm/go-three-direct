package handler

import (
	"github.com/idprm/go-three-direct/internal/domain/model"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/services"
)

type DRHandler struct {
	logger              *logger.Logger
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	req                 *model.DRRequest
}

func NewDRHandler(
	logger *logger.Logger,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	req *model.DRRequest,
) *DRHandler {
	return &DRHandler{
		logger:              logger,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		req:                 req,
	}
}

func (h *DRHandler) Push() error {
	return nil
}

// /**
//  * Sample Request
//  * {"msisdn":"62895635121559","shortcode":"99879","status":"DELIVRD","message":"1601666764269215859","ip":"116.206.10.222"}
//  */

// // parsing string json
// var req dto.DRRequest
// json.Unmarshal(message, &req)

// var transaction entity.Transaction
// existTrans := p.gdb.Where("msisdn", req.Msisdn).Where("submited_id", req.Message).First(&transaction)

// if existTrans.RowsAffected == 1 {

// 	var labelStatus string
// 	if req.Status == "DELIVRD" {
// 		labelStatus = "SUCCESS"
// 	} else {
// 		labelStatus = "FAILED"
// 	}

// 	transaction.Status = labelStatus
// 	transaction.DrStatus = req.Status
// 	transaction.DrStatusDetail = util.DRStatus(req.Status)
// 	p.gdb.Save(&transaction)
// }
