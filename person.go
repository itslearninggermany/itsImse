package itsImse

type Person struct {
	SyncPersonKey     string
	FirstName         string
	LastName          string
	NickName          string
	Birthday          string
	Profile           string
	Email             string
	Username          string
	Password          string
	Phone             string
	Mobile            string
	Address1          string
	Address2          string
	Postcode          string
	City              string
	CustomInformation map[string]string
	Child             []string
}
