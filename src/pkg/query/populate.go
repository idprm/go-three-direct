package query

import (
	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func GetDataPopulate(name string) ([]model.Subscription, error) {

	var SQL string

	switch name {
	case "RENEWAL":
		SQL = `SELECT id, msisdn, service_id, channel, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_trial = false AND is_active = true AND deleted_at IS null`
	case "RETRY":
		SQL = `SELECT id, msisdn, service_id, channel, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE_SUB(DATE(renewal_at), INTERVAL 1 DAY) = DATE(NOW()) AND is_trial = false AND is_retry = true AND is_active = true AND deleted_at IS null`
	case "PRERENEWAL":
		SQL = `SELECT id, msisdn, service_id, channel, ip_address FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE_SUB(DATE(renewal_at), INTERVAL 2 DAY) = DATE(NOW()) AND is_trial = false AND is_retry = false AND is_active = true AND deleted_at IS null`
	case "TEST":
		// SQL = `SELECT id, msisdn, service_id, channel, ip_address FROM subscriptions`
	}

	rows, err := database.Datasource.SqlDB().Query(SQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []model.Subscription

	for rows.Next() {

		var s model.Subscription
		if err := rows.Scan(&s.ID, &s.Msisdn, &s.ServiceID, &s.Channel, &s.IpAddress); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	if err = rows.Err(); err != nil {
		return subs, err
	}

	return subs, nil
}
