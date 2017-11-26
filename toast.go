package wns

import "encoding/xml"

type baseNotification struct {
	XMLName xml.Name
	Type    string `xml:"-"`
	Visual  visual `xml:"visual"`
	Audio   *audio `xml:"audio,omitempty"`
}

type visual struct {
	XMLName xml.Name `xml:"visual"`
	Binding binding  `xml:"binding"`
}

type audio struct {
	XMLName xml.Name `xml:"audio"`
	Src     string   `xml:"src,omitempty,attr"`
	Silent  bool     `xml:"silent,omitempty,attr"`
	Loop    bool     `xml:"loop,omitempty,attr"`
}

type binding struct {
	XMLName     xml.Name `xml:"binding"`
	Template    string   `xml:"template,attr"`
	BindingList []bindingElement
}

type bindingElement struct {
	XMLName xml.Name
	Id      string `xml:"id,attr"`
	Content string `xml:",innerxml"`
	Src     string `xml:"src,omitempty,attr"`
	Alt     string `xml:"alt,omitempty,attr"`
}

func (n *baseNotification) SetTemplate(template string) *baseNotification {
	n.Visual.Binding.Template = template
	return n
}

func (n *baseNotification) AddText(id string, content string) *baseNotification {
	newElement := bindingElement{Id: id, Content: content}
	newElement.XMLName.Local = "text"
	n.addBindingElement(newElement)
	return n
}

func (n *baseNotification) AddImage(id string, src string, alt string) *baseNotification {
	newElement := bindingElement{Id: id, Src: src, Alt: alt}
	newElement.XMLName.Local = "image"
	n.addBindingElement(newElement)
	return n
}

func (n *baseNotification) AddAudio(src string, silent bool, loop bool) {
	n.Audio = &audio{Src: src, Silent: silent, Loop: loop}
}

func (n *baseNotification) GetXml() (string, error) {
	return GetXml(n)
}

func (n *baseNotification) GetWnsType() (string) {
	return "wns/" + n.Type
}

func (n *baseNotification) addBindingElement(el bindingElement) {
	n.Visual.Binding.BindingList = append(n.Visual.Binding.BindingList, el)
}
