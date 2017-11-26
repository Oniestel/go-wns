package wns

import "encoding/xml"

type badge struct {
	XMLName xml.Name `xml:"badge"`
	Type    string   `xml:"-"`
	Value   string   `xml:"value"`
}

func (b *badge) SetValue(value string) *badge {
	b.Value = value
	return b
}

func (b *badge) GetXml() (string, error) {
	return GetXml(b)
}

func (b *badge) GetWnsType() (string) {
	return "wns/" + b.Type
}
