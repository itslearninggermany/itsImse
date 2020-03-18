package itsImse

import "encoding/xml"

type body struct {
	XMLName xml.Name `xml:"x:Body"`
	Content interface{}
}
