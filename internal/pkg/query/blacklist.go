package query

import (
	"database/sql"
)

type BlacklistRepository struct {
	db *sql.DB
}

type IBlacklistRepository interface {
	GetCountBlacklist(string) (int, error)
}

func NewBlacklistRepository(db *sql.DB) *BlacklistRepository {
	return &BlacklistRepository{
		db: db,
	}
}

func (r *BlacklistRepository) GetCountBlacklist(msisdn string) (int, error) {
	var count int
	sqlStatement := "SELECT COUNT(*) as count FROM blacklists WHERE msisdn = ? LIMIT 1"
	err := r.db.QueryRow(sqlStatement, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}
