package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"

	"waki.mobi/go-yatta-h3i/src/config"
)

func MessageTerminated() (string, error) {
	url := config.ViperEnv("URL_MT")

	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json; charset=utf8")

	if err != nil {
		return "", errors.New(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", errors.New(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string([]byte(body)), nil
}
