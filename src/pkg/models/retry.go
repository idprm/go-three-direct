package models

import "gorm.io/gorm"

type Retry struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Msisdn     string `gorm:"size:20" json:"msisdn"`
	SubmitId   string `gorm:"size:45" json:"submit_id"`
	Message    string `gorm:"150" json:"message"`
	Keyword    string `gorm:"100" json:"keyword"`
	SubKeyword string `gorm:"50" json:"sub_keyword"`
	RetryTo    uint   `gorm:"size:3" json:"retry_to"`
	gorm.Model
}
