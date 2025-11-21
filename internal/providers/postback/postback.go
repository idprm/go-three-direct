package postback

import (
	"log"

	"github.com/idprm/go-three-direct/internal/logger"
)

type Postback struct {
	logger *logger.Logger
}

func NewPostback(logger *logger.Logger) *Postback {
	return &Postback{
		logger: logger,
	}
}

type IPostback interface {
	Handle()
}

func (p *Postback) Handle() {
	log.Println("Handling postback...")
}
