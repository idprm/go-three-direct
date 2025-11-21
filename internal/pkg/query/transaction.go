package query

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/idprm/go-three-direct/internal/domain/entity"
)

type TransactionRepository struct {
	db *sql.DB
}

type ITransactionRepository interface {
	RemoveTransact(entity.Transaction) error
	InsertTransact(entity.Transaction) error
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) RemoveTransact(t entity.Transaction) error {
	query := "DELETE FROM transactions WHERE service_id = ? AND msisdn = ? AND subject = ? AND status = ? AND DATE(created_at) = DATE(?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, query)
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

func (r *TransactionRepository) InsertTransact(t entity.Transaction) error {
	query := "INSERT INTO transactions(transaction_id, service_id, msisdn, submited_id, keyword, adnet, amount, status, status_code, status_detail, subject, ip_address, payload, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, query)
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
