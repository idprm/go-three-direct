package dto

type MORequest struct {
	MobileNo  string `query:"mobile_no" json:"mobile_no"`
	ShortCode string `query:"short_code" json:"short_code"`
	Message   string `query:"message" json:"message"`
	IpAddress string `query:"ip" json:"ip"`
}

type DRRequest struct {
	Msisdn    string `query:"msisdn" json:"msisdn"`
	ShortCode string `query:"shortcode" json:"shortcode"`
	Status    string `query:"status" json:"status"`
	Message   string `query:"message" json:"message"`
	IpAddress string `query:"ip" json:"ip"`
}
