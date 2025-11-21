package cmd

import (
	"encoding/json"
	"sync"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/domain/model"
	"github.com/idprm/go-three-direct/internal/domain/repository"
	"github.com/idprm/go-three-direct/internal/handler"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/services"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Processor struct {
	db     *gorm.DB
	rds    *redis.Client
	logger *logger.Logger
}

func NewProcessor(
	db *gorm.DB,
	rds *redis.Client,
	logger *logger.Logger,
) *Processor {
	return &Processor{
		db:     db,
		rds:    rds,
		logger: logger,
	}
}

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
)

func (p *Processor) MO(wg *sync.WaitGroup, message []byte) {
	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)

	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)

	// parsing json to string
	var req *model.MORequest
	json.Unmarshal(message, &req)

	h := handler.NewMOHandler(
		p.logger,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		req,
	)

	if err := h.Register(); err != nil {
		panic(err)
	}

	wg.Done()
}

func (p *Processor) DR(wg *sync.WaitGroup, message []byte) {

	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)

	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)

	// parsing json to string
	var req *model.DRRequest
	json.Unmarshal(message, &req)

	h := handler.NewDRHandler(
		p.logger,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		req,
	)

	if err := h.Push(); err != nil {
		panic(err)
	}

	wg.Done()
}

func (p *Processor) Renewal(wg *sync.WaitGroup, message []byte) {

	serviceRepo := repository.NewServiceRepository(p.db)
	serviceService := services.NewServiceService(serviceRepo)

	contentRepo := repository.NewContentRepository(p.db)
	contentService := services.NewContentService(contentRepo)

	subscriptionRepo := repository.NewSubscriptionRepository(p.db)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	transactionRepo := repository.NewTransactionRepository(p.db)
	transactionService := services.NewTransactionService(transactionRepo)

	// parsing json to string
	var sub *entity.Subscription
	json.Unmarshal(message, &sub)

	h := handler.NewRenewalHandler(
		p.logger,
		serviceService,
		contentService,
		subscriptionService,
		transactionService,
		sub,
	)

	if err := h.Dailypush(); err != nil {
		panic(err)
	}

	wg.Done()
}

func (p *Processor) Retry(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}

func (p *Processor) Notif(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}

func (p *Processor) Purge(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}
