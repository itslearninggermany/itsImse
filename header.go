package itsImse

import (
	"encoding/xml"
	"fmt"
)

type Header header

type header struct {
	SyncRequestHeaderInfo syncRequestHeaderInfo `xml:"ims1:syncRequestHeaderInfo"`
	Security              security              `xml:"wsse:Security"`
}

type syncRequestHeaderInfo struct {
	MessageIdentifier string `xml:"ims1:messageIdentifier"`
}

func SetHeader(security *security, messageIdentifier string) *header {
	out := new(header)
	out.Security = *security
	syncreq := new(syncRequestHeaderInfo)
	syncreq.MessageIdentifier = messageIdentifier
	return out
}

func (p *header) ParseToXml() string {
	b, err := xml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

/*
   <x:Header>
        <ims1:syncRequestHeaderInfo>
            <ims1:messageIdentifier>1234567890</ims1:messageIdentifier>
        </ims1:syncRequestHeaderInfo>
        <wsse:Security xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
            <wsse:UsernameToken>
                <wsse:Username>cbcf5dc1-f410-4aca-9643-5a9cc5ecb17f</wsse:Username>
                <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">e3NnmNVRU7iXvms</wsse:Password>
            </wsse:UsernameToken>
        </wsse:Security>
    </x:Header>
*/
