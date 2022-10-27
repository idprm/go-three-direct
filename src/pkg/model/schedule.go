package model

import (
	"time"
)

type Schedule struct {
	ID         int       `gorm:"primaryKey;size:3" json:"id"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	PublishAt  time.Time `gorm:"default:null" json:"publish_at"`
	UnLockedAt time.Time `gorm:"default:null" json:"unlocked_at"`
	Status     bool      `gorm:"type:boolean" json:"status"`
}
