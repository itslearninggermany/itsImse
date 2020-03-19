package itsImse

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type imsEsClient struct {
	Envelope envelope
	username string
	password string
	security *security
	method   string
}

func NewImseClient(username, password string) *imsEsClient {
	p := new(imsEsClient)
	p.security = SetSecurity(username, password)
	p.Envelope = envelope{
		XMLName: xml.Name{},
		X:       x,
		Ims:     ims,
		Ims2:    ims2,
		Ims3:    ims3,
		Ims1:    ims1,
	}
	return p
}

func (p *imsEsClient) ShowXML() string {
	b, err := xml.Marshal(p.Envelope)
	fmt.Println(err)
	return string(b)
}

/*
sasd
*/
func (p *imsEsClient) Call() (response string, messageIdentifyer string) {
	head := setHeader(p.security, messageIdentifyer)
	p.Envelope.Header = *head
	soapbyte, err := xml.Marshal(p)
	if err != nil {
		return
	}
	soap := string(soapbyte)

	fmt.Println("======================================")
	fmt.Println(soap)
	fmt.Println("======================================")
	httpMethod := "POST"
	req, err := http.NewRequest(httpMethod, url, strings.NewReader(soap))
	if err != nil {
		return
	}

	soapAction, err := getSoapAction(p.method)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-type", `text/xml; charset="utf-8"`)
	req.Header.Set("SOAPAction", soapAction)

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	response = string(bodyBytes)
	return response, messageIdentifyer
}
