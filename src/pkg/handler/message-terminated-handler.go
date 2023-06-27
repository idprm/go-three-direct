package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/config"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

type Telco struct {
	cfg *config.Secret
}

func NewTelco(cfg *config.Secret) *Telco {
	return &Telco{
		cfg: cfg,
	}
}

func (p *Telco) MessageTerminated(service model.Service, content model.Content, msisdn string, transaction string) ([]byte, error) {
	loggerMT := util.MakeLogger("mt", true)

	urlAPI := p.cfg.Telco.Url

	payload := url.Values{}
	payload.Add("USERNAME", service.AuthUser)
	payload.Add("PASSWORD", service.AuthPass)
	payload.Add("REG_DELIVERY", "1")
	payload.Add("ORIGIN_ADDR", content.OriginAddr)
	payload.Add("MOBILENO", msisdn)
	payload.Add("TYPE", "0")
	payload.Add("MESSAGE", content.Value)
	payload.Add("UDH", "0")

	req, err := http.NewRequest("GET", urlAPI+"/push"+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerMT.WithFields(logrus.Fields{
		"url":            urlAPI + "?" + payload.Encode(),
		"msisdn":         msisdn,
		"transaction_id": transaction,
	}).Info("REQUEST")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return []byte(body), nil
}

func (p *Telco) MessageTerminatedUnknown(content model.Content, msisdn string, transaction string) ([]byte, error) {
	loggerMT := util.MakeLogger("mt", true)

	urlAPI := p.cfg.Telco.Url

	payload := url.Values{}
	payload.Add("USERNAME", "SD_210906_0180")
	payload.Add("PASSWORD", "y4tt43r4")
	payload.Add("REG_DELIVERY", "1")
	payload.Add("ORIGIN_ADDR", content.OriginAddr)
	payload.Add("MOBILENO", msisdn)
	payload.Add("TYPE", "0")
	payload.Add("MESSAGE", content.Value)
	payload.Add("UDH", "0")

	req, err := http.NewRequest("GET", urlAPI+"/push"+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerMT.WithFields(logrus.Fields{
		"url":            urlAPI + "?" + payload.Encode(),
		"msisdn":         msisdn,
		"transaction_id": transaction,
	}).Info("REQUEST")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return []byte(body), nil
}

func (p *Telco) MessageTerminatedRenewal(service model.Service, content model.Content, msisdn string, transaction string) ([]byte, error) {
	loggerMT := util.MakeLogger("mt", true)

	urlAPI := p.cfg.Telco.Url

	payload := url.Values{}
	payload.Add("USERNAME", service.AuthUser)
	payload.Add("PASSWORD", service.AuthPass)
	payload.Add("REG_DELIVERY", "1")
	payload.Add("ORIGIN_ADDR", content.OriginAddr)
	payload.Add("MOBILENO", msisdn)
	payload.Add("TYPE", "0")
	payload.Add("MESSAGE", content.Value)
	payload.Add("UDH", "0")

	req, err := http.NewRequest("GET", urlAPI+"/push"+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerMT.WithFields(logrus.Fields{
		"url":            urlAPI + "?" + payload.Encode(),
		"msisdn":         msisdn,
		"transaction_id": transaction,
	}).Info("REQUEST")

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: tr,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return []byte(body), nil
}
