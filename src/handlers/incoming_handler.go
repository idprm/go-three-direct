package handlers

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/config"
)

type IncomingHandler struct {
	cfg *config.Secret
	db  *sql.DB
}

func NewIncomingHandler(cfg *config.Secret, db *sql.DB) *IncomingHandler {
	return &IncomingHandler{
		cfg: cfg,
		db:  db,
	}
}

func (h *IncomingHandler) Sync() {

}
