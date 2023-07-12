package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/domain/entity"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

/**
 * NOTIF SUBSCRIPTION
 * Method: GET
 * Endpoint:
 */
func NotifSub(service entity.Service, msisdn string, transaction string) ([]byte, error) {
	loggerNotif := util.MakeLogger("notif", true)

	urlAPI := service.UrlNotifSub

	payload := url.Values{}
	payload.Add("msisdn", msisdn)
	payload.Add("package", "daily")
	payload.Add("event", "reg")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerNotif.WithFields(logrus.Fields{
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

/**
 * NOTIF UNSUBCRIPTION
 * Method: GET
 * Endpoint:
 */
func NotifUnsub(service entity.Service, msisdn string, transaction string) ([]byte, error) {
	loggerNotif := util.MakeLogger("notif", true)

	urlAPI := service.UrlNotifUnsub

	payload := url.Values{}
	payload.Add("msisdn", msisdn)
	payload.Add("event", "unreg")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerNotif.WithFields(logrus.Fields{
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

/**
 * NOTIF RENEWAL
 * Method: GET
 * Endpoint:
 */
func NotifRenewal(service entity.Service, msisdn string, transaction string) ([]byte, error) {
	loggerNotif := util.MakeLogger("notif", true)

	urlAPI := service.UrlNotifRenewal

	payload := url.Values{}
	payload.Add("msisdn", msisdn)
	payload.Add("package", "daily")
	payload.Add("event", "renewal")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerNotif.WithFields(logrus.Fields{
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
