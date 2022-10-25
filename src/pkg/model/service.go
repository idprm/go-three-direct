package model

type Service struct {
	ID              int    `gorm:"primaryKey;size:3" json:"id"`
	ServiceCode     string `gorm:"size:10" json:"service_code"`
	Category        string `gorm:"size:20" json:"category"`
	AuthUser        string `gorm:"size:45" json:"auth_user"`
	AuthPass        string `gorm:"size:45" json:"auth_pass"`
	Name            string `gorm:"size:250;not null" json:"name"`
	Day             int    `gorm:"size:3" json:"day"`
	Charge          int    `gorm:"size:6" json:"charge"`
	TrialDay        int    `gorm:"size:3" json:"trial_day"`
	UrlNotifSub     string `gorm:"size:150" json:"url_notif_sub"`
	UrlNotifUnsub   string `gorm:"size:150" json:"url_notif_unsub"`
	UrlNotifRenewal string `gorm:"size:150" json:"url_notif_renewal"`
	UrlPostback     string `gorm:"size:150" json:"url_postback"`
	IsActive        bool   `gorm:"type:boolean" json:"is_active"`
}
