package itsImse

type Person struct {
	syncPersonKey     string
	firstName         string
	lastName          string
	nickName          string
	birthday          string
	profile           string
	email             string
	username          string
	password          string
	phone             string
	mobile            string
	address1          string
	address2          string
	postcode          string
	city              string
	customInformation map[string]string
	child             []string
}
