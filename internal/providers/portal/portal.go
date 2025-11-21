package portal

import "github.com/idprm/go-three-direct/internal/logger"

type Portal struct {
	logger *logger.Logger
}

func NewPortal(logger *logger.Logger) *Portal {
	return &Portal{
		logger: logger,
	}
}

type IPortal interface {
	Handle()
}

func (p *Portal) Handle() error {
	return nil
}
