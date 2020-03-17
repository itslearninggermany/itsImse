package itsImse

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type client struct {
	username          string
	password          string
	messageIdentifyer string
	security          *security
	envelopeBody      interface{}
}

func NewClient(username, password string) *client {
	p := new(client)
	p.security = SetSecurity(username, password)
	return p
}

func (p *client) ShowRequestBody() string {
	p.messageIdentifyer = fmt.Sprint(uuid.Must(uuid.NewV4()))
	head := setHeader(p.security, p.messageIdentifyer)
	return head.parseToXml()
}

func (p *client) Call() (response string, messageIdentifyer string) {
	return "", ""
}
