package telco

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/idprm/go-three-direct/internal/domain/entity"
	"github.com/idprm/go-three-direct/internal/logger"
	"github.com/idprm/go-three-direct/internal/pkg/util"
	"github.com/idprm/go-three-direct/internal/utils"
	"github.com/sirupsen/logrus"
)

type Telco struct {
	logger  *logger.Logger
	service *entity.Service
	content *entity.Content
	sub     *entity.Subscription
}

func NewTelco(
	logger *logger.Logger,
	service *entity.Service,
	content *entity.Content,
	sub *entity.Subscription,
) *Telco {
	return &Telco{
		logger:  logger,
		service: service,
		content: content,
		sub:     sub,
	}
}

type ITelco interface {
	MobileTerminated() ([]byte, error)
}

func (p *Telco) MobileTerminated() ([]byte, error) {
	l := util.MakeLogger("mt", true)

	start := time.Now()
	trxId := utils.GenerateTrxId()

	req, err := http.NewRequest(http.MethodGet, p.service.GetUrlTelco(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	l.WithFields(logrus.Fields{
		"url_request":    p.service.GetUrlTelco(),
		"msisdn":         p.sub.GetMsisdn(),
		"transaction_id": trxId,
	}).Info("REQUEST_MT")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		DisableKeepAlives:  true,
	}

	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: tr,
	}
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	duration := time.Since(start).Milliseconds()
	p.logger.Writer(string(body))

	l.WithFields(logrus.Fields{
		"response":    string(body),
		"trx_id":      trxId,
		"duration":    duration,
		"status_code": resp.StatusCode,
		"status_text": http.StatusText(resp.StatusCode),
	}).Info("RESPONSE_MT")

	return body, nil
}
