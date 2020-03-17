package itsImse

import (
	"encoding/xml"
	"fmt"
)

const wsse = `http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd`
const wsu = `http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd`
const passwordtype = `http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText`

type security struct {
	Wsse          string `xml:"wsse,attr"`
	Wsu           string `xml:"wsu,attr"`
	UsernameToken struct {
		Username string `xml:"Username"`
		Password struct {
			Text string `xml:",chardata"`
			Type string `xml:"Type,attr"`
		} `xml:"Password"`
	} `xml:"UsernameToken"`
}

func SetSecurity(username, password string) *security {
	out := new(security)
	out.Wsse = wsse
	out.Wsu = wsu
	out.UsernameToken.Username = username
	out.UsernameToken.Password.Text = password
	out.UsernameToken.Password.Type = passwordtype
	return out
}

func (p *security) ParseToXml() string {
	b, err := xml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
