package handler

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/config"
)

type RetryHandler struct {
	cfg *config.Secret
	db  *sql.DB
}

func NewRetryHandler(cfg *config.Secret, db *sql.DB) *RetryHandler {
	return &RetryHandler{
		cfg: cfg,
		db:  db,
	}
}

func (h *RetryHandler) Firstpush() {

}

func (h *RetryHandler) Dailypush() {

}
