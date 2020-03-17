package itsImse

type imsePerson struct {
	FormatName formatName `xml:"formatName"`
	Name name `xml:"name"`
	Email  string `xml:"email"`
	UserId userID `xml:"userId"`
	Address address `xml:"address"`
	Demographics demographics `xml:"demographics"`
	InstitutionRole instituionRole `xml:"institutionRole"`
	Tel []telephone `xml:"tel"`
	Extension extension  `xml:"extension"`
	Relationship []relationship `xml:"relationship"`
}


type relationship struct {
	Relation string `xml:"relation"`
	SourceId sourcedId `xml:"sourceId"`
}

type sourcedId struct {
	Identifier string `xml:"identifier"`
}

type instituionRole struct {
	InstitutionRoleType string `xml:"institutionRoleType"`
	PrimaryRoleType     string `xml:"primaryRoleType"`
}

type name struct {
	PartName []partName `xml:"partname"`
}

type partName struct {
	NamePartType  string `xml:"namePartType"`
	NamePartValue string `xml:"namePartValue"`
}

type userID struct {
	UserIdValue string `xml:"userIdValue"`
	PassWord    string `xml:"passWord"`
}

type address struct {
	Locality string   `xml:"locality"`
	Postcode string   `xml:"postcode"`
	Country  string   `xml:"country"`
	Street   []string `xml:"street"`
}

type formatName struct {
	Nil  string `xml:"nil,attr"`
	Xsi  string `xml:"xsi,attr"`
}

type extension struct {
	ExtensionField [] extensionField `xml:"extensionField"`
}

type demographics struct {
	Bday string `xml:"bday"`
}

type extensionField struct {
	FieldName  string `xml:"fieldName"`
	FieldType  string `xml:"fieldType"`
	FieldValue string `xml:"fieldValue"`
}


type telephone struct {
	TelType  string `xml:"telType"`
	TelValue string `xml:"telValue"`
}

