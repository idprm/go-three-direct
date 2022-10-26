package dto

type Response struct {
	Body ResponseBodyXML `xml:"RESPONSES"`
}

type ResponseBodyXML struct {
	Param ResponseParamXML `xml:"RESPONSE"`
}

type ResponseParamXML struct {
	Code       int    `xml:"CODE"`
	Text       string `xml:"TEXT"`
	SubmitedID string `xml:"SUBMITTED_ID"`
}
