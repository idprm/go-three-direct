package entity

import (
	"time"
)

type Schedule struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	PublishAt  time.Time `json:"publish_at"`
	UnLockedAt time.Time `json:"un_locked_at"`
	Status     bool      `json:"status"`
}

func (e *Schedule) GetId() int {
	return e.ID
}

func (e *Schedule) GetName() string {
	return e.Name
}

func (e *Schedule) GetPublishAt() time.Time {
	return e.PublishAt
}

func (e *Schedule) GetUnlockedAt() time.Time {
	return e.UnLockedAt
}

func (e *Schedule) GetStatus() bool {
	return e.Status
}
