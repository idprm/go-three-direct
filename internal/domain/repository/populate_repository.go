package repository

import (
	"database/sql"

	"github.com/idprm/go-three-direct/internal/domain/entity"
)

const (
	querySelectSubRenewal = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
	querySelectSubRetry   = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE_SUB(DATE(renewal_at), INTERVAL 1 DAY) = DATE(NOW()) AND is_retry = true AND is_active = true AND deleted_at IS null ORDER BY success DESC"
)

type PopulateRepository struct {
	db *sql.DB
}

func NewPopulateRepository(db *sql.DB) *PopulateRepository {
	return &PopulateRepository{
		db: db,
	}
}

type IPopulateRepository interface {
	SelectRenewal() (*[]entity.Subscription, error)
	SelectRetry() (*[]entity.Subscription, error)
}

func (r *PopulateRepository) SelectRenewal() (*[]entity.Subscription, error) {
	rows, err := r.db.Query(querySelectSubRenewal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subs []entity.Subscription

	for rows.Next() {
		var s entity.Subscription
		if err := rows.Scan(&s.ID, &s.Msisdn, &s.ServiceID, &s.Keyword, &s.PurgeAt, &s.IpAddress, &s.CreatedAt); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return &subs, err
	}

	return &subs, nil
}

func (r *PopulateRepository) SelectRetry() (*[]entity.Subscription, error) {
	rows, err := r.db.Query(querySelectSubRetry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var subs []entity.Subscription

	for rows.Next() {
		var s entity.Subscription
		if err := rows.Scan(&s.ID, &s.Msisdn, &s.ServiceID, &s.Keyword, &s.PurgeAt, &s.IpAddress, &s.CreatedAt); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return &subs, err
	}

	return &subs, nil
}
