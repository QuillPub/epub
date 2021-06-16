package packagefile

import "encoding/xml"

// Meta contains meta information
// See https://www.w3.org/publishing/epub3/epub-packages.html#sec-meta-elem
type Meta struct {
	// Required
	Text     string
	Property Property

	//Optional
	Refines       Refines
	Scheme        Scheme
	XMLLang       XMLLang
	TextDirection TextDirection
	ID            ID
}

func (meta Meta) optionalAttributes() []optionalAttribute {
	return []optionalAttribute{
		meta.Refines,
		meta.Scheme,
		meta.XMLLang,
		meta.TextDirection,
		meta.ID,
	}
}

func (meta *Meta) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case meta.Property.Name():
			meta.Property = Property(attr.Value)
		case meta.Refines.Name():
			meta.Refines = Refines(attr.Value)
		case meta.TextDirection.Name():
			meta.TextDirection = TextDirection(attr.Value)
		case meta.Scheme.Name():
			meta.Scheme = Scheme(attr.Value)
		case meta.ID.Name():
			meta.ID = ID(attr.Value)
		case meta.XMLLang.Name():
			if attr.Name.Space != xmlNamespace {
				continue
			}
			meta.XMLLang = XMLLang(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.CharData:
			meta.Text = string(el)
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

// MarshalXML will not write out a meta element without any text
func (meta *Meta) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if meta.Text == "" {
		return nil
	}
	start.Name.Local = "meta"
	start.Attr = []xml.Attr{meta.Property.toAttr()}
	for _, attr := range meta.optionalAttributes() {
		if attr.isSet() {
			start.Attr = append(start.Attr, attr.toAttr())
		}
	}
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.CharData(meta.Text))
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}

// Property is the property specified by the meta tag, for example
// "dcterms:modified" specifies the last modified date for the document.
type Property string

func (property Property) String() string {
	return string(property)
}

// Name gives the xml attribute name
func (Property) Name() string {
	return "property"
}

func (property Property) xmlName() xml.Name {
	return xml.Name{Local: property.Name()}
}

func (property Property) isSet() bool {
	return string(property) != ""
}

func (property Property) toAttr() xml.Attr {
	return xml.Attr{Name: property.xmlName(), Value: property.String()}
}

// Refines links the meta tag to another element by specifying an element id
// in the format "#element-id"
type Refines string

func (refines Refines) String() string {
	return string(refines)
}

// Name gives the xml attribute name
func (Refines) Name() string {
	return "refines"
}

func (refines Refines) xmlName() xml.Name {
	return xml.Name{Local: refines.Name()}
}

func (refines Refines) isSet() bool {
	return string(refines) != ""
}

func (refines Refines) toAttr() xml.Attr {
	return xml.Attr{Name: refines.xmlName(), Value: refines.String()}
}

// Scheme identifies the scheme the value is drawn from
type Scheme string

func (scheme Scheme) String() string {
	return string(scheme)
}

// Name gives the xml attribute name
func (Scheme) Name() string {
	return "scheme"
}

func (scheme Scheme) xmlName() xml.Name {
	return xml.Name{Local: scheme.Name()}
}

func (scheme Scheme) isSet() bool {
	return string(scheme) != ""
}

func (scheme Scheme) toAttr() xml.Attr {
	return xml.Attr{Name: scheme.xmlName(), Value: scheme.String()}
}
