package query

import (
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func GetDataPopulate(name string) ([]model.Subscription, error) {

	var SQL string

	switch name {
	case "RENEWAL":
		SQL = `SELECT id, msisdn, service_id, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null`
	case "RETRY":
		SQL = `SELECT id, msisdn, service_id, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE_SUB(DATE(renewal_at), INTERVAL 1 DAY) = DATE(NOW()) AND is_retry = true AND is_active = true AND deleted_at IS null`
	case "TEST":
		// SQL = `SELECT id, msisdn, service_id, ip_address FROM subscriptions`
	}

	rows, err := database.Datasource.SqlDB().Query(SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []model.Subscription

	for rows.Next() {

		var s model.Subscription
		if err := rows.Scan(&s.ID, &s.Msisdn, &s.ServiceID, &s.IpAddress); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return subs, err
	}

	return subs, nil
}
