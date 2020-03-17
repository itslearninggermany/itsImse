package itsImse

import "encoding/xml"

const x = `http://schemas.xmlsoap.org/soap/envelope/`
const ims = `http://www.imsglobal.org/services/pms/xsd/imsPersonManMessSchema_v1p0`
const ims2 = `http://www.imsglobal.org/services/common/imsCommonSchema_v1p0`
const ims3 = `http://www.imsglobal.org/services/pms/xsd/imsPersonManDataSchema_v1p0`
const ims1 = `http://www.imsglobal.org/services/common/imsMessBindSchema_v1p0`

type envelope struct {
	XMLName xml.Name    `xml:"x:Envelope"`
	X       string      `xml:"xmlns:x,attr"`
	Ims     string      `xml:"xmlns:ims,attr"`
	Ims2    string      `xml:"xmlns:ims2,attr"`
	Ims3    string      `xml:"xmlns:ims3,attr"`
	Ims1    string      `xml:"xmlns:ims1,attr"`
	Header  header      `xml:"x:Header"`
	Body    interface{} `xml:"x:Body"`
}

func setEnvelope(header *header, body interface{}) *envelope {
	out := new(envelope)
	out.X = x
	out.Ims = ims
	out.Ims2 = ims2
	out.Ims3 = ims3
	out.Ims1 = ims1
	out.Header = *header
	out.Body = body
	return out
}
