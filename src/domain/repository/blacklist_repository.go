package repository

import (
	"database/sql"
)

const (
	queryCountBlacklist = "SELECT COUNT(*) as count FROM blacklists WHERE msisdn = ? LIMIT 1"
)

type BlacklistRepository struct {
	db *sql.DB
}

func NewBlacklistRepository(db *sql.DB) *BlacklistRepository {
	return &BlacklistRepository{
		db: db,
	}
}

type IBlacklistRepository interface {
	CountByMsisdn(string) (int, error)
}

func (r *BlacklistRepository) CountByMsisdn(msisdn string) (int, error) {
	var count int

	err := r.db.QueryRow(queryCountBlacklist, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}
