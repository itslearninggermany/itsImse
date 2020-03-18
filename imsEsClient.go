package itsImse

import (
	"encoding/xml"
	"fmt"
)

type imsEsClient struct {
	Envelope envelope
	username string
	password string
	security *security
}

func NewImseClient(username, password string) *imsEsClient {
	p := new(imsEsClient)
	p.security = SetSecurity(username, password)
	p.Envelope = envelope{
		XMLName: xml.Name{},
		X:       x,
		Ims:     ims,
		Ims2:    ims2,
		Ims3:    ims3,
		Ims1:    ims1,
	}
	return p
}

func (p *imsEsClient) ShowXML() string {
	b, err := xml.Marshal(p.Envelope)
	fmt.Println(err)
	return string(b)
}

/*
sasd
*/
func (p *imsEsClient) Call() (response string, messageIdentifyer string) {
	head := setHeader(p.security, messageIdentifyer)
	p.Envelope.Header = *head
	return "", ""
}
