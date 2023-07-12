package repository

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/domain/entity"
)

const (
	querySelectByName = "SELECT id, name, value FROM adnets WHERE name = ? LIMIT 1"
)

type AdnetRepository struct {
	db *sql.DB
}

func NewAdnetRepository(db *sql.DB) *BlacklistRepository {
	return &BlacklistRepository{
		db: db,
	}
}

type IAdnetRepository interface {
	GetAdnetByName(string) (*entity.Adnet, error)
}

func (r *AdnetRepository) GetAdnetByName(name string) (*entity.Adnet, error) {
	var a entity.Adnet
	err := r.db.QueryRow(querySelectByName, name).Scan(&a.ID, &a.Name, &a.Value)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
