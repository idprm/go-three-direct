package query

import (
	"database/sql"

	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

type ContentRepository struct {
	db *sql.DB
}

type IContentRepository interface {
	GetContent(int, string) (model.Content, error)
}

func NewContentRepository(db *sql.DB) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

func (r *ContentRepository) GetContent(serviceId int, name string) (model.Content, error) {
	var content model.Content
	sqlStatement := "SELECT value, origin_addr FROM contents WHERE service_id = ? AND name = ? LIMIT 1"
	err := r.db.QueryRow(sqlStatement, serviceId, name).Scan(&content.Value, &content.OriginAddr)
	if err != nil {
		return content, err
	}
	return content, nil
}
