package handlers

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/config"
)

type RenewalHandler struct {
	cfg *config.Secret
	db  *sql.DB
}

func NewRenewalHandler(cfg *config.Secret, db *sql.DB) *RenewalHandler {
	return &RenewalHandler{
		cfg: cfg,
		db:  db,
	}
}

func (h *RenewalHandler) Dailypush() {

}
