package handlers

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/config"
)

type DRHandler struct {
	cfg *config.Secret
	db  *sql.DB
}

func NewDRHandler(cfg *config.Secret, db *sql.DB) *DRHandler {
	return &DRHandler{
		cfg: cfg,
		db:  db,
	}
}

func (h *DRHandler) Sync() {

}
