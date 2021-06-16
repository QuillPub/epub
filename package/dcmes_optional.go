package packagefile

import "encoding/xml"

type Creator struct {
	DCMESOptionalElement
}
type Publisher struct {
	DCMESOptionalElement
}
type Description struct {
	DCMESOptionalElement
}
type Contributor struct {
	DCMESOptionalElement
}
type Coverage struct {
	DCMESOptionalElement
}
type Relation struct {
	DCMESOptionalElement
}
type Rights struct {
	DCMESOptionalElement
}
type Subject struct {
	DCMESOptionalElement
}
type Type struct {
	DCMESOptionalElement
}
type Source struct {
	DCMESOptionalElement
}

// DCMESOptionalElement is a template for several elements that are identical
// except for their names. This was the best solution I could come up with.
// See https://www.w3.org/publishing/epub3/epub-packages.html#sec-opf-dcmes-optional
type DCMESOptionalElement struct {
	Text string

	// Optional
	ID   ID
	Dir  TextDirection
	Lang XMLLang

	// name is the element name for the parent element and is used for
	// marshaling and unmarshaling xml
	name string
}

func (element DCMESOptionalElement) qualifiedName() string {
	return "dc:" + element.name
}

func (element DCMESOptionalElement) optionalAttributes() []optionalAttribute {
	return []optionalAttribute{element.Dir, element.ID, element.Lang}
}

func (element *DCMESOptionalElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case element.ID.Name():
			element.ID = ID(attr.Value)
		case element.Dir.Name():
			element.Dir = TextDirection(attr.Value)
		case element.Lang.Name():
			element.Lang = XMLLang(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.CharData:
			element.Text = string(el)
		case xml.StartElement:
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

// MarshalXML will not write a DCMESOptionalElement without any inner text.
func (element DCMESOptionalElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if element.Text == "" {
		return nil
	}
	start.Name.Local = element.qualifiedName()
	for _, attr := range element.optionalAttributes() {
		if attr.isSet() {
			start.Attr = append(start.Attr, attr.toAttr())
		}
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.CharData(element.Text))
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}
