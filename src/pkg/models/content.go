package models

type Content struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:45" json:"name"`
	Value string `gorm:"size:300" json:"value"`
}
