package model

type Content struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"size:45" json:"name"`
	OriginAddr string `gorm:"size:10" json:"origin_addr"`
	Value      string `gorm:"size:300" json:"value"`
}
