package dto

import "encoding/xml"

type Response struct {
	XMLName xml.Name        `xml:"RESPONSES"`
	Body    ResponseBodyXML `xml:"RESPONSE"`
}

type ResponseBodyXML struct {
	Code       int    `xml:"CODE"`
	Text       string `xml:"TEXT"`
	SubmitedID string `xml:"SUBMITTED_ID"`
}

type ResponseXML struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status"`
}
