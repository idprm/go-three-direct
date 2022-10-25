package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

const (
	detailNotifSub     = "NOTIF_SUBSCRIPTION"
	detailNotifUnsub   = "NOTIF_UNSUB"
	detailNotifRenewal = "NOTIF_RENEWAL"
	detailNone         = "-"
)

/**
 * NOTIF SUBSCRIPTION
 * Method: GET
 * Endpoint:
 */
func NotifSub(service model.Service, msisdn string, transaction string) ([]byte, error) {
	loggerNotif := util.MakeLogger("notif", true)

	urlAPI := service.UrlNotifSub

	payload := url.Values{}
	payload.Add("msisdn", msisdn)
	payload.Add("trxid", transaction)

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerNotif.WithFields(logrus.Fields{
		"request_url": urlAPI + "?" + payload.Encode(),
		"msisdn":      msisdn,
	}).Info(detailNotifSub)

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
func NotifUnsub(service model.Service, msisdn string) ([]byte, error) {
	loggerNotif := util.MakeLogger("notif", true)

	urlAPI := service.UrlNotifUnsub

	payload := url.Values{}
	payload.Add("msisdn", msisdn)

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerNotif.WithFields(logrus.Fields{
		"request_url": urlAPI + "?" + payload.Encode(),
		"msisdn":      msisdn,
	}).Info(detailNotifUnsub)

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
func NotifRenewal(service model.Service, msisdn string, transaction string) ([]byte, error) {
	loggerNotif := util.MakeLogger("notif", true)

	urlAPI := service.UrlNotifRenewal

	payload := url.Values{}
	payload.Add("msisdn", msisdn)
	payload.Add("trxid", transaction)

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	loggerNotif.WithFields(logrus.Fields{
		"request_url": urlAPI + "?" + payload.Encode(),
		"msisdn":      msisdn,
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
