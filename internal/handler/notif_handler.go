package handler

import (
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/services"
)

type NotifHandler struct {
	logger              *logger.Logger
	serviceService      services.IServiceService
	subscriptionService services.ISubscriptionService
}

func NewNotifHandler(
	logger *logger.Logger,
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
) *NotifHandler {
	return &NotifHandler{
		logger:              logger,
		serviceService:      serviceService,
		subscriptionService: subscriptionService,
	}
}

func (h *NotifHandler) Sub() error {
	return nil
}

func (h *NotifHandler) Unsub() error {
	return nil
}

func (h *NotifHandler) Renewal() error {
	return nil
}
