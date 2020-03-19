package itsImse

import "errors"

const url string = `https://enterprise.itsltest.com/WCFServiceLibrary/ImsEnterpriseServicesPort.svc`

func getSoapAction(input string) (out string, err error) {
	switch input {
	case "createPerson":
		return "http://www.imsglobal.org/soap/pms/createPerson", nil

	default:
		return "", errors.New("No method in struct")
	}
}
