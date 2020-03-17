package itsImse

type Header header

type header struct {
	SyncRequestHeaderInfo syncRequestHeaderInfo `xml:"syncRequestHeaderInfo"`
	Security security `xml:"Security"`
}

type syncRequestHeaderInfo struct {
	MessageIdentifier string `xml:"messageIdentifier"`
}


