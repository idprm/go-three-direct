package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint64 `gorm:"primaryKey" json:"id"`
	TransactionID string `gorm:"size:45" json:"transaction_id"`
	Msisdn        string `gorm:"size:25" json:"msisdn"`
	ServiceID     int    `gorm:"size:3" json:"-"`
	Service       Service
	SubmitedID    string  `gorm:"size:50" json:"submited_id"`
	Keyword       string  `gorm:"size:50" json:"keyword"`
	Adnet         string  `gorm:"size:55;default:null" json:"adnet"`
	Amount        float64 `gorm:"size:6;default:0" json:"amount"`
	Status        string  `gorm:"size:45;default:null" json:"status"`
	StatusDetail  string  `gorm:"size:45;default:null" json:"status_detail"`
	Subject       string  `gorm:"size:45;default:null" json:"subject"`
	IpAddress     string  `gorm:"size:45;default:null" json:"ip_address"`
	Payload       string  `gorm:"type:text" json:"payload"`
	gorm.Model
}
