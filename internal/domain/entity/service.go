package entity

type Service struct {
	ID              int     `json:"id"`
	Code            string  `json:"code"`
	Name            string  `json:"name"`
	AuthUser        string  `json:"auth_user"`
	AuthPass        string  `json:"auth_pass"`
	Day             int     `json:"day"`
	Charge          float64 `json:"charge"`
	PurgeDay        int     `json:"purge_day"`
	UrlNotifSub     string  `json:"url_notif_sub"`
	UrlNotifUnsub   string  `json:"url_notif_unsub"`
	UrlNotifRenewal string  `json:"url_notif_renewal"`
	UrlPostback     string  `json:"url_postback"`
	UrlTelco        string  `json:"url_telco"`
	IsActive        bool    `json:"is_active"`
}

func (e *Service) GetId() int {
	return e.ID
}

func (e *Service) GetCode() string {
	return e.Code
}

func (e *Service) GetAuthUser() string {
	return e.AuthUser
}

func (e *Service) GetAuthPass() string {
	return e.AuthPass
}

func (e *Service) GetDay() int {
	return e.Day
}

func (e *Service) GetCharge() float64 {
	return e.Charge
}

func (e *Service) GetPurgeDay() int {
	return e.PurgeDay
}

func (e *Service) GetUrlNotifSub() string {
	return e.UrlNotifSub
}

func (e *Service) GetUrlNotifUnsub() string {
	return e.UrlNotifUnsub
}

func (e *Service) GetUrlNotifRenewal() string {
	return e.UrlNotifRenewal
}

func (e *Service) GetUrlPostback() string {
	return e.UrlPostback
}

func (e *Service) GetIsActive() bool {
	return e.IsActive
}

func (e *Service) GetUrlTelco() string {
	return e.UrlTelco
}

// q.Add("ORIGIN_ADDR", p.content.GetOriginAddr())
// q.Add("MOBILENO", p.sub.GetMsisdn())
// q.Add("MESSAGE", p.content.GetValue())
