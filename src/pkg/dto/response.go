package dto

import "encoding/xml"

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

type ResponseXML struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status"`
}
