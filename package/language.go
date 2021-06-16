package packagefile

import (
	"encoding/xml"
)

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
		case xml.StartElement:
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

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
