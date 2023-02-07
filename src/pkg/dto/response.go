package dto

import "encoding/xml"

type Response struct {
	XMLName xml.Name        `xml:"RESPONSES" json:"RESPONSES"`
	Body    ResponseBodyXML `xml:"RESPONSE" json:"RESPONSE"`
}

type ResponseBodyXML struct {
	Code       int    `xml:"CODE" json:"CODE"`
	Text       string `xml:"TEXT" json:"TEXT"`
	SubmitedID string `xml:"SUBMITTED_ID" json:"SUBMITTED_ID"`
}

type ResponseXML struct {
	XMLName xml.Name `xml:"response" json:"response"`
	Status  string   `xml:"status" json:"status"`
}
