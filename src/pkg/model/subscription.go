package model

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	ID            uint64 `gorm:"primaryKey" json:"id"`
	ServiceID     int    `gorm:"size:3" json:"service_id"`
	Service       Service
	SubmitedID    string    `gorm:"size:50" json:"submited_id"`
	Msisdn        string    `gorm:"size:25" json:"msisdn"`
	Keyword       string    `gorm:"100" json:"keyword"`
	Adnet         string    `gorm:"size:55;default:null" json:"adnet"`
	LatestSubject string    `gorm:"size:45;default:null" json:"latest_subject"`
	LatestStatus  string    `gorm:"size:45;default:null" json:"latest_status"`
	Amount        float64   `gorm:"size:6;default:0" json:"amount"`
	RenewalAt     time.Time `gorm:"default:null" json:"renewal_at"`
	PurgeAt       time.Time `gorm:"default:null" json:"purge_at"`
	UnsubAt       time.Time `gorm:"default:null" json:"unsub_at"`
	ChargeAt      time.Time `gorm:"default:null" json:"charge_at"`
	RetryAt       time.Time `gorm:"default:null" json:"retry_at"`
	Success       uint      `gorm:"size:4;default:null" json:"success"`
	IpAddress     string    `gorm:"size:45;default:null" json:"ip_address"`
	IsRetry       bool      `gorm:"type:bool" json:"is_retry"`
	IsActive      bool      `gorm:"type:bool" json:"is_active"`
	gorm.Model
}
