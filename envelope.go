package itsImse

type envelope struct {
	Header header      `xml:"x:Header"`
	Body   interface{} `xml:"x:Body"`
}

func setEnvelope(header *header, body interface{}) *envelope {
	out := new(envelope)
	out.Header = *header
	out.Body = body
	return out
}
