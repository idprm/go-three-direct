package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	TransactionID  string `gorm:"size:45" json:"transaction_id"`
	ServiceID      int    `gorm:"size:3" json:"-"`
	Service        Service
	Msisdn         string  `gorm:"size:25" json:"msisdn"`
	SubmitedID     string  `gorm:"size:50" json:"submited_id"`
	Keyword        string  `gorm:"size:50" json:"keyword"`
	Adnet          string  `gorm:"size:55;default:null" json:"adnet"`
	Amount         float64 `gorm:"size:6;default:0" json:"amount"`
	Status         string  `gorm:"size:45;default:null" json:"status"`
	StatusCode     int     `gorm:"size:5;default:null" json:"status_code"`
	StatusDetail   string  `gorm:"size:100;default:null" json:"status_detail"`
	Subject        string  `gorm:"size:45;default:null" json:"subject"`
	DrStatus       string  `gorm:"size:45;default:null" json:"dr_status"`
	DrStatusDetail string  `gorm:"size:100;default:null" json:"dr_status_detail"`
	IpAddress      string  `gorm:"size:45;default:null" json:"ip_address"`
	Payload        string  `gorm:"type:text" json:"payload"`
	gorm.Model
}
