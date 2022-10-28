package model

type Service struct {
	ID              int     `gorm:"primaryKey;size:3" json:"id"`
	Code            string  `gorm:"size:15" json:"code"`
	Name            string  `gorm:"size:50;not null" json:"name"`
	AuthUser        string  `gorm:"size:45" json:"auth_user"`
	AuthPass        string  `gorm:"size:45" json:"auth_pass"`
	Day             int     `gorm:"size:3" json:"day"`
	Charge          float64 `gorm:"size:6" json:"charge"`
	PurgeDay        int     `gorm:"size:3" json:"purge_day"`
	UrlNotifSub     string  `gorm:"size:150" json:"url_notif_sub"`
	UrlNotifUnsub   string  `gorm:"size:150" json:"url_notif_unsub"`
	UrlNotifRenewal string  `gorm:"size:150" json:"url_notif_renewal"`
	UrlPostback     string  `gorm:"size:150" json:"url_postback"`
	IsActive        bool    `gorm:"type:boolean" json:"is_active"`
}
