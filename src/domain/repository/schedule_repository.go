package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"waki.mobi/go-yatta-h3i/src/domain/entity"
)

const (
	querySchedulePublishAt = "SELECT COUNT(*) as count FROM schedules WHERE name = ? AND TIME(publish_at) = ? AND status = true"
	queryScheduleLockedAt  = "SELECT COUNT(*) as count FROM schedules WHERE name = ? AND TIME(un_locked_at) = ? AND status = false"
	queryScheduleUpdate    = "UPDATE schedules SET status = ? WHERE name = ?"
)

type ScheduleRepository struct {
	db *sql.DB
}

func NewScheduleRepository(db *sql.DB) *ScheduleRepository {
	return &ScheduleRepository{
		db: db,
	}
}

type IScheduleRepository interface {
	CountUnlocked(string, string) (int, error)
	CountLocked(string, string) (int, error)
	ScheduleUpdate(*entity.Schedule) error
}

func (r *ScheduleRepository) CountUnlocked(name string, time string) (int, error) {
	var count int
	err := r.db.QueryRow(querySchedulePublishAt, name, time).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ScheduleRepository) CountLocked(name string, time string) (int, error) {
	var count int
	err := r.db.QueryRow(queryScheduleLockedAt, name, time).Scan(&count)
	if err != nil {
		return count, err
	}
	return count, nil
}

func (r *ScheduleRepository) ScheduleUpdate(s *entity.Schedule) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := r.db.PrepareContext(ctx, queryScheduleUpdate)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, s.Status, s.Name)
	if err != nil {
		log.Printf("Error %s when update row into schedules table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d schedules updated ", rows)

	return nil
}
