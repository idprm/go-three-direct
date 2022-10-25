package cmd

import "sync"

func moProccesor(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}

func drProccesor(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}

func renewalProccesor(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}

func retryProccesor(wg *sync.WaitGroup, message []byte) {

	wg.Done()
}
