package models

type Service struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:45" json:"name"`
	Code        string `gorm:"size:45" json:"code"`
	UrlPostback string `gorm:"size:100" json:"url_postback"`
	UrlNotif    string `gorm:"size:100" json:"url_notif"`
}
