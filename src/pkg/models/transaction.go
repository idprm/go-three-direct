package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID          uint64 `gorm:"primaryKey" json:"id"`
	Msisdn      string `gorm:"size:20" json:"msisdn"`
	Keyword     string `gorm:"100" json:"keyword"`
	SubKeyword  string `gorm:"50" json:"sub_keyword"`
	Data        string `gorm:"150" json:"data"`
	SubmittedId string `gorm:"200" json:"submited_id"`
	gorm.Model
}
