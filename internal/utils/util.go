package utils

import (
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

func GenerateTrxId() string {
	id := uuid.New()
	return id.String()
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Panicf("Error %v", key)
	}
	return strings.TrimSpace(value)
}
