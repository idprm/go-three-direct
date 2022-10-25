package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/pkg/config"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

func MessageTerminated(service model.Service, content model.Content, msisdn string, transaction string) ([]byte, error) {
	loggerMT := util.MakeLogger("mt", true)

	urlAPI := config.ViperEnv("URL_MT")

	payload := url.Values{}
	payload.Add("username", service.AuthUser)
	payload.Add("password", service.AuthPass)
	payload.Add("reg_delivery", "1")
	payload.Add("origin_addr", content.OriginAddr)
	payload.Add("mobileno", msisdn)
	payload.Add("type", "0")
	payload.Add("message", content.Value)
	payload.Add("udh", "0")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerMT.WithFields(logrus.Fields{
		"request_url":    urlAPI + "?" + payload.Encode(),
		"msisdn":         msisdn,
		"transaction_id": transaction,
	}).Info(detailNotifRenewal)

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
