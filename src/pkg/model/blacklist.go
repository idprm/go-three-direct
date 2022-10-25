package model

import "gorm.io/gorm"

type Blacklist struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Msisdn string `gorm:"size:25" json:"msisdn"`
	gorm.Model
}
