package itsImse

const wsse = `http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd`
const wsu = `http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd`
const passwordtype = `http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-Username-token-Profile-1.0#PasswordText`

type security struct {
	Wsse          string `xml:"xmlns:wsse,attr"`
	Wsu           string `xml:"xmlns:wsu,attr"`
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
