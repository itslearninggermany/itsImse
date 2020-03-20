package itsImse

import "encoding/xml"

type readPersonsRequest struct {
	XMLName      xml.Name `xml:"ims:readPersonsRequest"`
	SourcedIdSet sourcedIdSet
}

type sourcedIdSet struct {
	XMLName    xml.Name `xml:"ims:sourcedIdSet"`
	Identifier string   `xml:"ims2:identifier"`
}

func (p *imsEsClient) ReadPerson(identifier string) *imsEsClient {
	p.Envelope.Body = body{
		Content: readPersonsRequest{
			SourcedIdSet: sourcedIdSet{
				Identifier: identifier,
			},
		},
	}
	p.method = "readPerson"
	return p
}
