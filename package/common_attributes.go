package packagefile

import "encoding/xml"

const xmlNamespace = `http://www.w3.org/XML/1998/namespace`
const dcNamespace = `http://purl.org/dc/elements/1.1/`

type optionalAttribute interface {
	isSet() bool
	toAttr() xml.Attr
}

type TextDirection string

const (
	TextDirectionLTR TextDirection = `ltr`
	TextDirectionRTL TextDirection = `rtl`
)

func (d TextDirection) String() string {
	return string(d)
}

func (d TextDirection) Name() string {
	return "dir"
}

func (d TextDirection) xmlName() xml.Name {
	return xml.Name{Local: d.Name()}
}

func (d TextDirection) isSet() bool {
	return string(d) != ""
}

func (d TextDirection) toAttr() xml.Attr {
	return xml.Attr{Name: d.xmlName(), Value: d.String()}
}

type ID string

func (id ID) String() string {
	return string(id)
}

func (ID) Name() string {
	return "id"
}

func (id ID) xmlName() xml.Name {
	return xml.Name{Local: id.Name()}
}

func (id ID) isSet() bool {
	return string(id) != ""
}

func (id ID) toAttr() xml.Attr {
	return xml.Attr{Name: id.xmlName(), Value: id.String()}
}

type XMLLang string

func (lang XMLLang) String() string {
	return string(lang)
}

func (XMLLang) Name() string {
	return "lang"
}

func (lang XMLLang) xmlName() xml.Name {
	return xml.Name{Local: "xml:" + lang.Name()}
}

func (lang XMLLang) isSet() bool {
	return string(lang) != ""
}

func (lang XMLLang) toAttr() xml.Attr {
	return xml.Attr{Name: lang.xmlName(), Value: lang.String()}
}
