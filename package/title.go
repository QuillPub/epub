package packagefile

import "encoding/xml"

// Title identifies a title of the document. At least one Title is required. If
// a title and subtitle are used, meta tags should identify which is which. The
// first Title is considered the primary title.
type Title struct {
	// Required
	Text string

	// Optional
	ID   ID
	Dir  TextDirection
	Lang XMLLang
}

func (title Title) optionalAttributes() []optionalAttribute {
	return []optionalAttribute{title.Dir, title.ID, title.Lang}
}

func (title *Title) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case title.Dir.Name():
			title.Dir = TextDirection(attr.Value)
		case title.ID.Name():
			title.ID = ID(attr.Value)
		case title.Lang.Name():
			title.Lang = XMLLang(attr.Value)
		}

	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.CharData:
			title.Text = string(el)
		case xml.StartElement:
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

// MarshalXML will not write a Title without any Text
func (title Title) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if title.Text == "" {
		return nil
	}
	start.Name.Local = "dc:title"
	for _, attr := range title.optionalAttributes() {
		if attr.isSet() {
			start.Attr = append(start.Attr, attr.toAttr())
		}
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.CharData(title.Text))
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}
