package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

func Postback(service model.Service, msisdn string, adnet string, transaction string) ([]byte, error) {
	loggerPb := util.MakeLogger("postback", true)

	urlAPI := service.UrlPostback

	payload := url.Values{}
	payload.Add("partner", "yatta-h3i")
	payload.Add("px", adnet)
	payload.Add("serv_id", strings.ToLower(service.Name))
	payload.Add("msisdn", msisdn)
	payload.Add("trxid", transaction)

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerPb.WithFields(logrus.Fields{
		"url":         urlAPI + "?" + payload.Encode(),
		"msisdn":      msisdn,
		"transaction": transaction,
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
