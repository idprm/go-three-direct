package handlers

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/config"
)

type MOHandler struct {
	cfg *config.Secret
	db  *sql.DB
}

func NewMOHandler(cfg *config.Secret, db *sql.DB) *MOHandler {
	return &MOHandler{
		cfg: cfg,
		db:  db,
	}
}

func (h *MOHandler) Firstpush() {

}
