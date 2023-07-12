package repository

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/domain/entity"
)

const (
	queryGetContent = "SELECT value, origin_addr FROM contents WHERE service_id = ? AND name = ? LIMIT 1"
)

type ContentRepository struct {
	db *sql.DB
}

func NewContentRepository(db *sql.DB) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

type IContentRepository interface {
	Get(int, string) (*entity.Content, error)
}

func (r *ContentRepository) Get(serviceId int, name string) (*entity.Content, error) {
	var content entity.Content
	err := r.db.QueryRow(queryGetContent, serviceId, name).Scan(&content.Value, &content.OriginAddr)
	if err != nil {
		return nil, err
	}
	return &content, nil
}
