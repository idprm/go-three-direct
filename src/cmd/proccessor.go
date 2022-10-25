package cmd

import (
	"encoding/json"
	"log"
	"sync"

	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func moProccesor(wg *sync.WaitGroup, message []byte) {
	// parsing string json
	var sub model.Subscription
	json.Unmarshal(message, &sub)

	log.Println(sub.Msisdn)

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
