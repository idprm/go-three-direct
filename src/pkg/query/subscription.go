package query

import (
	"context"
	"database/sql"
	"log"
	"time"

	"waki.mobi/go-yatta-h3i/src/database"
	"waki.mobi/go-yatta-h3i/src/pkg/model"
)

func GetSub(serviceId int, msisdn string) (model.Subscription, error) {
	var sub model.Subscription
	sqlStatement := `SELECT service_id, msisdn FROM subscriptions WHERE service_id = ? AND msisdn = ? LIMIT 1`

	db := database.Datasource.SqlDB()
	err := db.QueryRow(sqlStatement, serviceId, msisdn).Scan(&sub.ServiceID, &sub.Msisdn)
	if err != nil {
		return sub, err
	}

	return sub, nil
}

func SubUpdateLatest(db *sql.DB, s model.Subscription) error {
	query := "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.LatestSubject, s.LatestStatus, s.ServiceID, s.Msisdn)
	if err != nil {
		log.Printf("Error %s when update row into subscriptions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d subscriptions updated ", rows)

	return nil
}

func SubUpdateSuccess(db *sql.DB, s model.Subscription) error {
	query := "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, amount = amount + ?, renewal_at = ?, charge_at = ?, success = success + ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.LatestSubject, s.LatestStatus, s.Amount, s.RenewalAt, s.ChargeAt, s.Success, s.IsRetry, s.ServiceID, s.Msisdn)
	if err != nil {
		log.Printf("Error %s when update row into subscriptions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d subscriptions updated ", rows)

	return nil
}

func SubUpdateFailed(db *sql.DB, s model.Subscription) error {
	query := "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, renewal_at = ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.LatestSubject, s.LatestStatus, s.RenewalAt, s.IsRetry, s.ServiceID, s.Msisdn)
	if err != nil {
		log.Printf("Error %s when update row into subscriptions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d subscriptions updated ", rows)

	return nil
}
