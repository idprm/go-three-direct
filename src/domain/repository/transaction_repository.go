package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"waki.mobi/go-yatta-h3i/src/domain/entity"
)

const (
	queryDeleteTrans    = "DELETE FROM transactions WHERE service_id = ? AND msisdn = ? AND subject = ? AND status = ? AND DATE(created_at) = DATE(?)"
	queryInsertTrans    = "INSERT INTO transactions(transaction_id, service_id, msisdn, submited_id, keyword, adnet, amount, status, status_code, status_detail, subject, ip_address, payload, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	queryUpdateTrans    = "UPDATE transactions SET status = ?, status_code = ?, status_detail = ?, subject = ?, dr_status= ?, dr_status_detail = ?, updated_at = NOW() WHERE service_id = ? AND msisdn = ?"
	queryUpdateSubmited = "UPDATE transactions SET status = ?, dr_status = ?, dr_status_detail = ?, updated_at = NOW() WHERE submited_id = ? AND msisdn = ?"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

type ITransactionRepository interface {
	RemoveTransact(*entity.Transaction) error
	InsertTransact(*entity.Transaction) error
	UpdateTransact(*entity.Transaction) error
	UpdateSubmitedTransact(*entity.Transaction) error
}

func (r *TransactionRepository) RemoveTransact(t *entity.Transaction) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryDeleteTrans)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, t.ServiceID, t.Msisdn, t.Subject, t.Status, time.Now())
	if err != nil {
		log.Printf("Error %s when remove row into transactions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d transactions deleted ", rows)
	return nil
}

func (r *TransactionRepository) InsertTransact(t *entity.Transaction) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryInsertTrans)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, t.TransactionID, t.ServiceID, t.Msisdn, t.SubmitedID, t.Keyword, t.Adnet, t.Amount, t.Status, t.StatusCode, t.StatusDetail, t.Subject, t.IpAddress, t.Payload, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error %s when inserting row into transactions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d transactions created ", rows)
	return nil
}

func (r *TransactionRepository) UpdateTransact(t *entity.Transaction) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryUpdateTrans)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, t.Status, t.StatusCode, t.StatusDetail, t.Subject, t.DrStatus, t.DrStatusDetail, t.ServiceID, t.Msisdn)
	if err != nil {
		log.Printf("Error %s when update row into transactions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d transactions updated ", rows)

	return nil
}

func (r *TransactionRepository) UpdateSubmitedTransact(t *entity.Transaction) error {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryUpdateSubmited)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, t.Status, t.DrStatus, t.DrStatusDetail, t.SubmitedID, t.Msisdn)
	if err != nil {
		log.Printf("Error %s when update row into transactions table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d transactions updated ", rows)

	return nil
}
