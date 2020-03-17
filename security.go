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
		Username string `xml:"wsse:Username"`
		Password struct {
			Text string `xml:",chardata"`
			Type string `xml:"Type,attr"`
		} `xml:"wsse:Password"`
	} `xml:"wsse:UsernameToken"`
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

/*
<security wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	<UsernameToken>
		<Username>tnordmann</Username>
		<Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">pat5rxe4</Password>
	</UsernameToken>
</security>
*/

/*

   <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
       <wsse:UsernameToken>
           <wsse:Username>cbcf5dc1-f410-4aca-9643-5a9cc5ecb17f</wsse:Username>
           <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">e3NnmNVRU7iXvms</wsse:Password>
       </wsse:UsernameToken>
   </wsse:Security>
*/
