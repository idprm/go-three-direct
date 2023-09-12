package handler

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"waki.mobi/go-yatta-h3i/src/config"
)

type MOHandler struct {
	cfg *config.Secret
}

func NewMOHandler(cfg *config.Secret) *MOHandler {
	return &MOHandler{
		cfg: cfg,
	}
}

func Firstpush(msisdn string) ([]byte, error) {
	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */
	urlAPI := "http://35.247.131.49/moh3i"

	payload := url.Values{}
	payload.Add("mobile_no", msisdn)
	payload.Add("short_code", "99879")
	payload.Add("message", "REG GMPN")

	req, err := http.NewRequest("GET", urlAPI+"?"+payload.Encode(), nil)
	if err != nil {
		return nil, errors.New(err.Error())
	}

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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return body, nil
}
