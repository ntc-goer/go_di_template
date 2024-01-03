package util

import "encoding/xml"

type SenderIdentifier struct {
	SenderIDType string `xml:"SenderIDType"`
	IDTypeName   string `xml:"IDTypeName"`
	IDValue      string `xml:"IDValue"`
}

type Sender struct {
	SenderIdentifier SenderIdentifier `xml:"SenderIdentifier"`
	SenderName       string           `xml:"SenderName"`
	ContactName      string           `xml:"ContactName"`
	EmailAddress     string           `xml:"EmailAddress"`
}

type Header struct {
	Sender       Sender `xml:"Sender"`
	SentDateTime string `xml:"SentDateTime"`
}

type ProductIdentifier struct {
	ProductIDType string `xml:"ProductIDType"`
	IDValue       string `xml:"IDValue"`
}

type Product struct {
	RecordReference   string            `xml:"RecordReference"`
	NotificationType  string            `xml:"NotificationType"`
	ProductIdentifier ProductIdentifier `xml:"ProductIdentifier"`
}

type XMLProduct struct {
	ONIXMessage xml.Name   `xml:"ONIXMessage"`
	Header      Header     `xml:"Header"`
	Product     []*Product `xml:"Product"`
}
