package dto

type DRRequest struct {
	Msisdn    string `form:"msisdn" json:"msisdn"`
	ShortCode string `form:"shortcode" json:"shortcode"`
	Status    string `form:"status" json:"status"`
	Message   string `form:"message" json:"message"`
	IpAddress string `form:"ip" json:"ip"`
}

type MORequest struct {
	MobileNo  string `form:"mobile_no" json:"mobile_no"`
	ShortCode string `form:"short_code" json:"short_code"`
	Message   string `form:"message" json:"message"`
	IpAddress string `form:"ip" json:"ip"`
}
