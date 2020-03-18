package itsImse

import "encoding/xml"

const x = `http://schemas.xmlsoap.org/soap/envelope/`
const ims = `http://www.imsglobal.org/services/pms/xsd/imsPersonManMessSchema_v1p0`
const ims2 = `http://www.imsglobal.org/services/common/imsCommonSchema_v1p0`
const ims3 = `http://www.imsglobal.org/services/pms/xsd/imsPersonManDataSchema_v1p0`
const ims1 = `http://www.imsglobal.org/services/common/imsMessBindSchema_v1p0`

type envelope struct {
	XMLName xml.Name `xml:"x:Envelope"`
	X       string   `xml:"xmlns:x,attr"`
	Ims     string   `xml:"xmlns:ims,attr"`
	Ims2    string   `xml:"xmlns:ims2,attr"`
	Ims3    string   `xml:"xmlns:ims3,attr"`
	Ims1    string   `xml:"xmlns:ims1,attr"`
	Header  header
	Body    body
}
