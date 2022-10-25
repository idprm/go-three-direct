package model

type Keyword struct {
	ID   int    `gorm:"primaryKey;size:3"`
	Name string `gorm:"size:10"`
}
