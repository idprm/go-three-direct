package query

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/idprm/go-three-direct/internal/domain/entity"
)

type SubscriptionRepository struct {
	db *sql.DB
}

type ISubscriptionRepository interface {
	GetSub(int, string) (entity.Subscription, error)
	SubUpdateLatest(entity.Subscription) error
	SubUpdateSuccess(entity.Subscription) error
	SubUpdateFailed(entity.Subscription)
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

func (r *SubscriptionRepository) GetSub(serviceId int, msisdn string) (entity.Subscription, error) {
	var s entity.Subscription
	sqlStatement := `SELECT id, service_id, msisdn, keyword, adnet, latest_subject, latest_status, amount, renewal_at, purge_at, unsub_at, charge_at, retry_at, success, ip_address, is_retry, is_purge, is_active, created_at, updated_at FROM subscriptions WHERE service_id = ? AND msisdn = ? AND deleted_at IS NULL LIMIT 1`
	err := r.db.QueryRow(sqlStatement, serviceId, msisdn).Scan(&s.ID, &s.ServiceID, &s.Msisdn, &s.Keyword, &s.Adnet, &s.LatestSubject, &s.LatestStatus, &s.Amount, &s.RenewalAt, &s.PurgeAt, &s.UnsubAt, &s.ChargeAt, &s.RetryAt, &s.Success, &s.IpAddress, &s.IsRetry, &s.IsPurge, &s.IsActive, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (r *SubscriptionRepository) SubUpdateLatest(s entity.Subscription) error {
	query := "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, query)
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

func (r *SubscriptionRepository) SubUpdateSuccess(s entity.Subscription) error {
	query := "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, amount = amount + ?, renewal_at = ?, charge_at = ?, success = success + ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, query)
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

func (r *SubscriptionRepository) SubUpdateFailed(s entity.Subscription) error {
	query := "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, renewal_at = ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, query)
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
