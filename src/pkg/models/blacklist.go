package models

import "gorm.io/gorm"

type Blacklist struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Msisdn string `gorm:"size:20" json:"msisdn"`
	gorm.Model
}
