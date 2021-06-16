package packagefile

import "encoding/xml"

const xmlNamespace = `http://www.w3.org/XML/1998/namespace`
const dcNamespace = `http://purl.org/dc/elements/1.1/`

type optionalAttribute interface {
	isSet() bool
	toAttr() xml.Attr
}

// TextDirection annotates the direction of the text in the contained content
// and attributes
type TextDirection string

const (
	TextDirectionLTR TextDirection = `ltr`
	TextDirectionRTL TextDirection = `rtl`
)

func (d TextDirection) String() string {
	return string(d)
}

// Name gives the xml attribute name
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

// ID is an identifier for the element. This will be referenced by other elements
// such as meta elements.
type ID string

func (id ID) String() string {
	return string(id)
}

// Name gives the xml attribute name
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

// XMLLang specifies the language of the contents
type XMLLang string

func (lang XMLLang) String() string {
	return string(lang)
}

// Name gives the xml attribute name
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
