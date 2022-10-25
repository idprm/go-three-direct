package cmd

import (
	"encoding/json"
	"log"
	"strings"
	"sync"

	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/dto"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
	"waki.mobi/go-yatta-h3i/src/pkg/util"
)

const (
	reg          = "REG"
	unreg        = "UNREG"
	welcome      = "WELCOME"
	registration = "REGISTRATION"
	confirmation = "CONFIRMATION"
	firstpush    = "FIRSTPUSH"
	renewal      = "RENEWAL"
	unsub        = "UNSUB"
	erroyKey     = "ERROR_KEYWORD"
	failed       = "FAILED"
	reminder     = "REMINDER"
	isActive     = "IS_ACTIVE"
	purge        = "PURGE"
)

func moProccesor(wg *sync.WaitGroup, message []byte) {

	/**
	 * {"mobile_no":"62895330590144","short_code":"99879","message":"REG KEREN","ip":"116.206.10.222"}
	 */
	transactionId := util.GenerateTransactionId()

	// init global variable
	var (
		keyword string
		subject string
		payload string
	)

	// parsing string json
	var req dto.MORequest
	json.Unmarshal(message, &req)

	log.Println(req.ShortCode)

	//
	var service model.Service
	database.Datasource.DB().Where("code", req.ShortCode).First(&service)

	// checking subscription
	var subscription model.Subscription
	activeSub := database.Datasource.DB().Where("msisdn", req.MobileNo).Where("is_active", true).First(&subscription)

	if activeSub.RowsAffected == 0 {
		database.Datasource.DB().Create(&model.Subscription{
			Msisdn:    req.MobileNo,
			ServiceID: service.ID,
			IpAddress: req.IpAddress,
			IsActive:  true,
		})
	}

	if activeSub.RowsAffected == 1 {
		subscription.IpAddress = req.IpAddress
		database.Datasource.DB().Save(&subscription)
	}

	// split message param
	msg := strings.Split(req.Message, " ")
	// define array with index
	index0 := strings.ToUpper(msg[0])
	index1 := strings.ToUpper(msg[1])
	// indexKeyword := ""

	if index0 != reg && index0 != unreg {
		subject = erroyKey
	}

	if index0 == reg {
		subject = firstpush
	}

	if index0 == unreg {
		subject = unreg
	}

	if index1 != service.Name {
		//
	}

	/**
	 * Query Content
	 */
	var contWelcome model.Content
	database.Datasource.DB().Where("name", welcome).First(&contWelcome)

	var contConfirm model.Content
	database.Datasource.DB().Where("name", confirmation).First(&contConfirm)

	var contWrongKey model.Content
	database.Datasource.DB().Where("name", erroyKey).First(&contWrongKey)

	var contUnsub model.Content
	database.Datasource.DB().Where("name", unsub).First(&contUnsub)

	var contIsActive model.Content
	database.Datasource.DB().Where("name", isActive).First(&contIsActive)

	database.Datasource.DB().Create(
		&model.Transaction{
			TransactionID: transactionId,
			ServiceID:     service.ID,
			Msisdn:        req.MobileNo,
			Keyword:       keyword,
			Amount:        0,
			Subject:       subject,
			IpAddress:     req.IpAddress,
			Payload:       payload,
		},
	)

	wg.Done()
}

func drProccesor(wg *sync.WaitGroup, message []byte) {
	// parsing string json
	var sub model.Subscription
	json.Unmarshal(message, &sub)

	log.Println(sub.Msisdn)

	wg.Done()
}

func renewalProccesor(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}

func retryProccesor(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}
