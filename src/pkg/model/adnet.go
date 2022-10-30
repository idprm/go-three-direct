package model

type Adnet struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"size:10" json:"name"`
	Value string `gorm:"size:10" json:"value"`
}
