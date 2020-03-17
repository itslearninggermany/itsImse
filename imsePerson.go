package itsImse

type imsePerson struct {
	FormatName      formatName     `xml:"ims3:formatName"`
	Name            name           `xml:"ims3:name"`
	Email           string         `xml:"ims2:email"`
	UserId          userID         `xml:"ims3:userId"`
	Address         address        `xml:"ims3:address"`
	Demographics    demographics   `xml:"ims3:demographics"`
	InstitutionRole instituionRole `xml:"ims3:institutionRole"`
	Tel             []telephone    `xml:"ims3:tel"`
	Extension       extension      `xml:"ims3:extension"`
	Relationship    []relationship `xml:"ims2:relationship"`
}

type relationship struct {
	Relation string    `xml:"ims2:relation"`
	SourceId sourcedId `xml:"ims2:sourceId"`
}

type sourcedId struct {
	Identifier string `xml:"ims2:identifier"`
}

type instituionRole struct {
	InstitutionRoleType string `xml:"ims3:institutionRoleType"`
	PrimaryRoleType     string `xml:"ims3:primaryRoleType"`
}

type name struct {
	PartName []partName `xml:"ims3:partname"`
}

type partName struct {
	NamePartType  string `xml:"ims3:namePartType"`
	NamePartValue string `xml:"ims3:namePartValue"`
}

type userID struct {
	UserIdValue string `xml:"ims2:userIdValue"`
	PassWord    string `xml:"ims2:passWord"`
}

type address struct {
	Locality string   `xml:"ims3:locality"`
	Postcode string   `xml:"ims3:postcode"`
	Street   []string `xml:"ims3:street"`
}

const xsinil = `true`
const xsi = `http://www.w3.org/2001/XMLSchema-instance`

type formatName struct {
	Nil string `xml:"xsi:nil,attr"`
	Xsi string `xml:"xmlns:xsi,attr"`
}

type extension struct {
	ExtensionField []extensionField `xml:"ims2:extensionField"`
}

type demographics struct {
	Bday string `xml:"ims3:bday"`
}

type extensionField struct {
	FieldName  string `xml:"ims2:fieldName"`
	FieldType  string `xml:"ims2:fieldType"`
	FieldValue string `xml:"ims2:fieldValue"`
}

type telephone struct {
	TelType  string `xml:"ims3:telType"`
	TelValue string `xml:"ims3:telValue"`
}

func createImsePerson(person Person) (x imsePerson, syncPersonKey string) {
	var relationships []relationship
	if len(person.child) != 0 {
		for i := 0; i < len(person.child); i++ {
			relationships = append(relationships, relationship{
				Relation: "Child",
				SourceId: sourcedId{Identifier: person.child[i]},
			})
		}
	} else {
		relationships = nil
	}

	var extensionFields []extensionField
	if len(person.customInformation) != 0 {
		for key, value := range person.customInformation {
			extensionFields = append(extensionFields, extensionField{
				FieldName:  key,
				FieldType:  "string",
				FieldValue: value,
			})
		}
	} else {
		extensionFields = nil
	}

	var telephones []telephone
	telephones = append(telephones, telephone{
		TelType:  "Mobile",
		TelValue: person.mobile,
	})
	telephones = append(telephones, telephone{
		TelType:  "Voice",
		TelValue: person.phone,
	})

	var streets []string
	streets = append(streets, person.address1)
	streets = append(streets, person.address2)

	var partNames []partName
	partNames = append(partNames, partName{
		NamePartType:  "First",
		NamePartValue: person.firstName,
	})
	partNames = append(partNames, partName{
		NamePartType:  "Last",
		NamePartValue: person.lastName,
	})
	partNames = append(partNames, partName{
		NamePartType:  "Nick",
		NamePartValue: person.nickName,
	})

	x = imsePerson{
		FormatName: formatName{
			Nil: xsinil,
			Xsi: xsi,
		},
		Name:  name{PartName: partNames},
		Email: person.email,
		UserId: userID{
			UserIdValue: person.syncPersonKey,
			PassWord:    person.password,
		},
		Address: address{
			Locality: person.city,
			Postcode: person.postcode,
			Street:   streets,
		},
		Demographics: demographics{Bday: person.birthday},
		InstitutionRole: instituionRole{
			InstitutionRoleType: person.profile,
			PrimaryRoleType:     "true",
		},
		Tel:          telephones,
		Extension:    extension{ExtensionField: extensionFields},
		Relationship: relationships,
	}

	return x, person.syncPersonKey
}
