package packagefile

import "encoding/xml"

// Identifier is some kind of identifier for the publication, for example a UUID
// or a publisher's id. If the identifier comes from an established system, an
// accompanying meta tag for the property "identifier-type" should be provided.
// See https://www.w3.org/publishing/epub3/epub-packages.html#sec-opf-dcidentifier
type Identifier struct {
	// Required
	Text string

	//Optional
	ID ID
}

func (identifier *Identifier) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == identifier.ID.Name() {
			identifier.ID = ID(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.CharData:
			identifier.Text = string(el)
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

// MarshalXML will not write an identifier without any Text
func (identifier Identifier) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if identifier.Text == "" {
		return nil
	}
	start.Name.Local = "dc:identifier"
	if identifier.ID.isSet() {
		start.Attr = []xml.Attr{identifier.ID.toAttr()}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.CharData(identifier.Text))
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}
