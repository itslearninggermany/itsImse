package itsImse

import (
	"encoding/xml"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type client struct {
	Envelope     envelope `xml:"x:Envelope"`
	username     string
	password     string
	security     *security
	envelopeBody *interface{}
}

func NewClient(username, password string) *client {
	p := new(client)
	p.security = SetSecurity(username, password)
	return p
}

func (p *client) ShowRequestBodyFromTheLastCall() string {
	b, err := xml.Marshal(p.Envelope)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

/*
sasd
*/
func (p *client) Call() (response string, messageIdentifyer string) {
	messageIdentifyer = fmt.Sprint(uuid.Must(uuid.NewV4()))
	head := setHeader(p.security, messageIdentifyer)
	p.Envelope = *setEnvelope(head, p.envelopeBody)
	return "", ""
}
