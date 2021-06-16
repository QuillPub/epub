package packagefile

import (
	"encoding/xml"
)

// Language designates a language used in the publication. Metadata must contain
// at least one Language. Multiple languages  may be specified in the case of
// multilingual publication, but the first Language is considered primary. Not
// that this does not specify the language of any child content.
// See https://www.w3.org/publishing/epub3/epub-packages.html#sec-opf-dclanguage
type Language struct {
	// Required
	Text string

	// Optional
	ID ID
}

func (language *Language) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == language.ID.Name() {
			language.ID = ID(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.CharData:
			language.Text = string(el)
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

// MarshalXML will not write a language without any Text
func (language Language) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if language.Text == "" {
		return nil
	}
	start.Name.Local = "dc:language"
	if language.ID.isSet() {
		start.Attr = []xml.Attr{language.ID.toAttr()}
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.CharData(language.Text))
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}
