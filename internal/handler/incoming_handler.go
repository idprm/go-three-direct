package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/services"
	"github.com/redis/go-redis/v9"
	"github.com/wiliehidayat87/rmqp"
)

const (
	SUBJECT_WELCOME   string = "WELCOME"
	SUBJECT_FIRSTPUSH string = "FIRSTPUSH"
	SUBJECT_RENEWAL   string = "RENEWAL"
	SUBJECT_UNSUB     string = "UNSUB"
	SUBJECT_INSUFT    string = "INSUFT"
	SUBJECT_ERROYKEY  string = "ERROR_KEYWORD"
	SUBJECT_ISACTIVE  string = "IS_ACTIVE"
	SUBJECT_PURGE     string = "PURGE"
	SUBJECT_UNKNOWN   string = "UNKNOWN_KEYWORD"

	SMS_FIRSTPUSH string = "MT_FIRSTPUSH"
	SMS_WELCOME   string = "MT_WELCOME"
	SMS_ISACTIVE  string = "MT_ISACTIVE"
	SMS_PURGE     string = "MT_PURGE"
	SMS_RENEWAL   string = "MT_RENEWAL"
	SMS_INSUFF    string = "MT_INSUFFICIENT"
	SMS_UNSUB     string = "MT_UNSUB"
	SMS_WRONGKEY  string = "MT_WRONGKEY"

	STATUS_SUCCESS string = "SUCCESS"
	STATUS_FAILED  string = "FAILED"
	STATUS_PENDING string = "PENDING"
)

type IncomingHandler struct {
	rds                 *redis.Client
	rmq                 rmqp.AMQP
	l                   *logger.Logger
	blacklistService    services.IBlacklistService
	serviceService      services.IServiceService
	contentService      services.IContentService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
}

func NewIncomingHandler(
	rds *redis.Client,
	rmq rmqp.AMQP,
	l *logger.Logger,
	blacklistService services.IBlacklistService,
	serviceService services.IServiceService,
	contentService services.IContentService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
) *IncomingHandler {
	return &IncomingHandler{
		rds:                 rds,
		rmq:                 rmq,
		l:                   l,
		blacklistService:    blacklistService,
		serviceService:      serviceService,
		contentService:      contentService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
	}
}

func (h *IncomingHandler) MobileOriginated(c *fiber.Ctx) error {
	return nil
}

func (h *IncomingHandler) DeliveryReport(c *fiber.Ctx) error {
	return nil
}
