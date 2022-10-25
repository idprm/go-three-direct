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
	Keyword       string  `gorm:"size:50" json:"keyword"`
	SubmittedId   string  `gorm:"size:20" json:"submited_id"`
	Channel       string  `gorm:"size:20" json:"channel"`
	Adnet         string  `gorm:"size:55;default:null" json:"adnet"`
	PubID         string  `gorm:"size:55;default:null" json:"pub_id"`
	AffSub        string  `gorm:"size:100;default:null" json:"aff_sub"`
	Amount        float64 `gorm:"size:6;default:0" json:"amount"`
	Status        string  `gorm:"size:45;default:null" json:"status"`
	StatusDetail  string  `gorm:"size:45;default:null" json:"status_detail"`
	Subject       string  `gorm:"size:45;default:null" json:"subject"`
	IpAddress     string  `gorm:"size:45;default:null" json:"ip_address"`
	Payload       string  `gorm:"type:text" json:"payload"`
	gorm.Model
}
