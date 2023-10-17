package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"waki.mobi/go-yatta-h3i/src/domain/entity"
)

const (
	queryGetSub                 = "SELECT id, service_id, msisdn, keyword, adnet, latest_subject, latest_status, amount, renewal_at, purge_at, unsub_at, charge_at, retry_at, success, ip_address, is_retry, is_purge, is_active, created_at, updated_at FROM subscriptions WHERE service_id = ? AND msisdn = ? AND deleted_at IS NULL LIMIT 1"
	queryCountRetrySub          = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = ? AND msisdn = ? AND is_retry = true AND is_active = true"
	queryCountSub               = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = ? AND msisdn = ?"
	queryCountNotActiveSub      = "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = ? AND msisdn = ? AND is_active = false"
	queryCountActiveSubByMsisdn = "SELECT COUNT(*) as count FROM subscriptions WHERE msisdn = ? AND is_active = true"
	queryUpdateKeyword          = "UPDATE subscriptions SET keyword = ?, adnet = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
	queryUpdateLatest           = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
	queryUpdateSuccess          = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, amount = amount + ?, renewal_at = ?, charge_at = ?, success = success + ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
	queryUpdateFailed           = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, renewal_at = ?, is_retry = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ? AND is_active = true"
	querySubUpdateDisable       = "UPDATE subscriptions SET latest_subject = ?, latest_status = ?, unsub_at = ?, is_retry = false, is_active = false, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
	queryInsertSub              = "INSERT INTO subscriptions(service_id, msisdn, keyword, adnet, latest_subject, latest_status, ip_address, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	queryRenewalAll             = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
	queryRenewalOdd             = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE MOD(CAST(RIGHT(msisdn, 1) AS INT), 2) = 1 AND renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
	queryRenewalEven            = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE MOD(CAST(RIGHT(msisdn, 1) AS INT), 2) = 0 AND renewal_at IS NOT NULL AND DATE(renewal_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
	queryRetry                  = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE_SUB(DATE(renewal_at), INTERVAL 1 DAY) = DATE(NOW()) AND is_retry = true AND is_active = true AND deleted_at IS null ORDER BY success DESC"
	queryPurge                  = "SELECT id, msisdn, service_id, keyword, purge_at, ip_address, created_at FROM subscriptions WHERE renewal_at IS NOT NULL AND DATE(purge_at) <= DATE(NOW()) AND is_active = true AND deleted_at IS null ORDER BY success DESC"
)

type SubscriptionRepository struct {
	db *sql.DB
}

func NewSubscriptionRepository(db *sql.DB) *SubscriptionRepository {
	return &SubscriptionRepository{
		db: db,
	}
}

type ISubscriptionRepository interface {
	GetSub(int, string) (*entity.Subscription, error)
	GetCountRetrySub(int, string) (int, error)
	GetCountSub(int, string) (int, error)
	GetCountActiveSub(int, string) (int, error)
	GetCountNotActiveSub(int, string) (int, error)
	SubUpdateKeyword(*entity.Subscription) error
	SubUpdateLatest(*entity.Subscription) error
	SubUpdateSuccess(*entity.Subscription) error
	SubUpdateFailed(*entity.Subscription) error
	SubUpdateDisable(*entity.Subscription) error
	InsertSub(*entity.Subscription) error
	GetRenewal() (*[]entity.Subscription, error)
	GetRetry() (*[]entity.Subscription, error)
}

func (r *SubscriptionRepository) GetSub(serviceId int, msisdn string) (*entity.Subscription, error) {
	var s entity.Subscription
	err := r.db.QueryRow(queryGetSub, serviceId, msisdn).Scan(&s.ID, &s.ServiceID, &s.Msisdn, &s.Keyword, &s.Adnet, &s.LatestSubject, &s.LatestStatus, &s.Amount, &s.RenewalAt, &s.PurgeAt, &s.UnsubAt, &s.ChargeAt, &s.RetryAt, &s.Success, &s.IpAddress, &s.IsRetry, &s.IsPurge, &s.IsActive)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *SubscriptionRepository) GetCountRetrySub(serviceId int, msisdn string) (int, error) {
	var count int
	err := r.db.QueryRow(queryCountRetrySub, serviceId, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) GetCountSub(serviceId int, msisdn string) (int, error) {
	var count int
	err := r.db.QueryRow(queryCountSub, serviceId, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) GetCountActiveSub(serviceId int, msisdn string) (int, error) {
	var count int
	sqlStatement := "SELECT COUNT(*) as count FROM subscriptions WHERE service_id = ? AND msisdn = ? AND is_active = true"
	err := r.db.QueryRow(sqlStatement, serviceId, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) GetCountNotActiveSub(serviceId int, msisdn string) (int, error) {
	var count int
	err := r.db.QueryRow(queryCountNotActiveSub, serviceId, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) GetCountActiveSubByMsisdn(msisdn string) (int, error) {
	var count int
	err := r.db.QueryRow(queryCountActiveSubByMsisdn, msisdn).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *SubscriptionRepository) SubUpdateKeyword(s *entity.Subscription) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryUpdateKeyword)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.Keyword, s.Adnet, s.ServiceID, s.Msisdn)
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

func (r *SubscriptionRepository) SubUpdateLatest(s *entity.Subscription) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryUpdateLatest)
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

func (r *SubscriptionRepository) SubUpdateSuccess(s *entity.Subscription) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryUpdateSuccess)
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

func (r *SubscriptionRepository) SubUpdateFailed(s *entity.Subscription) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryUpdateFailed)
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

func (r *SubscriptionRepository) SubUpdateDisable(s *entity.Subscription) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, querySubUpdateDisable)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.LatestSubject, s.LatestStatus, s.UnsubAt, s.ServiceID, s.Msisdn)
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

func (r *SubscriptionRepository) InsertSub(s *entity.Subscription) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryInsertSub)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.ServiceID, s.Msisdn, s.Keyword, s.Adnet, s.LatestSubject, s.LatestStatus, s.IpAddress, s.IsActive, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error %s when inserting row into subscriptions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d subscriptions created ", rows)
	return nil
}

func (r *SubscriptionRepository) GetRenewal() (*[]entity.Subscription, error) {
	rows, err := r.db.Query(queryRenewalAll)
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
		return nil, err
	}

	return &subs, nil
}

func (r *SubscriptionRepository) GetRetry() (*[]entity.Subscription, error) {
	rows, err := r.db.Query(queryRetry)
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
		return nil, err
	}

	return &subs, nil
}
