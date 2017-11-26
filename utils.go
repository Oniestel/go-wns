package wns

import "encoding/xml"

func GetXml(notification interface{}) (string, error) {
	xmlRes, err := xml.MarshalIndent(notification, "", "    ")
	if err == nil {
		xmlRes = []byte(xml.Header + string(xmlRes))
		return string(xmlRes), nil
	}
	return "", err
}
