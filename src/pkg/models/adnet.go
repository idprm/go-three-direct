package models

type Adnet struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:5" json:"name"`
	Value string `gorm:"size:5" json:"value"`
}
