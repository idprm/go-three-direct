package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Msisdn     string `gorm:"size:20" json:"msisdn"`
	Message    string `gorm:"150" json:"message"`
	Keyword    string `gorm:"100" json:"keyword"`
	SubKeyword string `gorm:"50" json:"sub_keyword"`
	IsActive   bool   `gorm:"type:bool" json:"is_active"`
	RenewalAt  *time.Time
	PurgeAt    *time.Time
	gorm.Model
}
