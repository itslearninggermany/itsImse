package itsImse

import "encoding/xml"

type createPersonRequest struct {
	XMLName   xml.Name   `xml:"ims:createPersonRequest"`
	SourcedId sourcedId  `xml:"ims:sourcedId"`
	Person    imsePerson `xml:"ims:person"`
}

func (p *imsEsClient) CreatePerson(person Person) *imsEsClient {
	imsePerson, syncPersonKey := createImsePerson(person)
	p.envelopeBody.request = createPersonRequest{
		SourcedId: sourcedId{Identifier: syncPersonKey},
		Person:    imsePerson,
	}
	return p
}
