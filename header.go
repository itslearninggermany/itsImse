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

func setHeader(security *security, messageIdentifier string) *header {
	out := new(header)
	out.Security = *security
	out.SyncRequestHeaderInfo = *setSyncRequestHeaderInfo(messageIdentifier)
	return out
}

func setSyncRequestHeaderInfo(messageIdentifier string) *syncRequestHeaderInfo {
	out := new(syncRequestHeaderInfo)
	out.MessageIdentifier = messageIdentifier
	return out
}

func (p *header) parseToXml() string {
	b, err := xml.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
