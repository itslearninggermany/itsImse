package itsImse

import (
	"encoding/xml"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type imsEsClient struct {
	Envelope     envelope
	username     string
	password     string
	security     *security
	envelopeBody *interface{} `xml:"x:Body"`
}

func NewImseClient(username, password string) *imsEsClient {
	p := new(imsEsClient)
	p.security = SetSecurity(username, password)
	return p
}

func (p *imsEsClient) ShowRequestBodyFromTheLastCall() string {
	b, err := xml.Marshal(p.Envelope)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func (p *imsEsClient) SetEnvelopeBody(in *interface{}) {
	p.envelopeBody = in
}

/*
sasd
*/
func (p *imsEsClient) Call() (response string, messageIdentifyer string) {
	messageIdentifyer = fmt.Sprint(uuid.Must(uuid.NewV4()))
	head := setHeader(p.security, messageIdentifyer)
	p.Envelope = *setEnvelope(head, p.envelopeBody)
	return "", ""
}
