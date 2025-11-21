package query

import (
	"database/sql"

	"github.com/idprm/go-three-direct/internal/domain/entity"
)

type PopulateRepository struct {
	db *sql.DB
}

type IPopulateRepository interface {
	GetDataPopulate(string) ([]entity.Subscription, error)
}

func NewPopulateRepository(db *sql.DB) *PopulateRepository {
	return &PopulateRepository{
		db: db,
	}
}

func (r *PopulateRepository) GetDataPopulate(name string) ([]entity.Subscription, error) {

	var SQL string

	switch name {
	case "RENEWAL":
		SQL = `SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY DATE(created_at) DESC, success DESC`
	case "RENEWAL_ODD":
		SQL = `SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE MOD(CAST(RIGHT(msisdn, 1) AS INT), 2) = 1 AND renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY DATE(created_at) DESC, success DESC`
	case "RENEWAL_EVEN":
		SQL = `SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE MOD(CAST(RIGHT(msisdn, 1) AS INT), 2) = 0 AND renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY DATE(created_at) DESC, success DESC`
	case "RETRY":
		SQL = `SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE_SUB(DATE(renewal_at), INTERVAL 1 DAY) = DATE(NOW()) AND is_retry = true AND is_active = true AND deleted_at IS null ORDER BY DATE(created_at) DESC, success DESC`
	case "PURGE":
		SQL = `SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(purge_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY DATE(created_at) DESC, success DESC`
	}

	rows, err := r.db.Query(SQL)
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
		return subs, err
	}

	return subs, nil
}
