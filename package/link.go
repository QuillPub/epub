package packagefile

import "encoding/xml"

type Link struct {
	// Required
	Href Href
	Rel  Rel

	// Required for external resources
	MediaType MediaType

	// Optional
	ID         ID
	Properties Properties
	Refines    Refines
}

func (link Link) optionalAttributes() []optionalAttribute {
	return []optionalAttribute{link.MediaType, link.ID, link.Properties, link.Refines}
}

func (link *Link) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case link.Href.Name():
			link.Href = Href(attr.Value)
		case link.Rel.Name():
			link.Rel = Rel(attr.Value)
		case link.MediaType.Name():
			link.MediaType = MediaType(attr.Value)
		case link.ID.Name():
			link.ID = ID(attr.Value)
		case link.Properties.Name():
			link.Properties = Properties(attr.Value)
		case link.Refines.Name():
			link.Refines = Refines(attr.Value)
		}

	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

func (link Link) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if !link.Href.isSet() || !link.Rel.isSet() {
		return nil
	}
	start.Name.Local = "link"
	start.Attr = []xml.Attr{link.Href.toAttr(), link.Rel.toAttr()}
	for _, attr := range link.optionalAttributes() {
		if attr.isSet() {
			start.Attr = append(start.Attr, attr.toAttr())
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}

type Href string

func (href Href) String() string {
	return string(href)
}

func (Href) Name() string {
	return "href"
}

func (href Href) xmlName() xml.Name {
	return xml.Name{Local: "xml:" + href.Name()}
}

func (href Href) isSet() bool {
	return string(href) != ""
}

func (href Href) toAttr() xml.Attr {
	return xml.Attr{Name: href.xmlName(), Value: href.String()}
}

type Rel string

func (rel Rel) String() string {
	return string(rel)
}

func (Rel) Name() string {
	return "rel"
}

func (rel Rel) xmlName() xml.Name {
	return xml.Name{Local: "xml:" + rel.Name()}
}

func (rel Rel) isSet() bool {
	return string(rel) != ""
}

func (rel Rel) toAttr() xml.Attr {
	return xml.Attr{Name: rel.xmlName(), Value: rel.String()}
}

type MediaType string

func (mediaType MediaType) String() string {
	return string(mediaType)
}

func (MediaType) Name() string {
	return "media-type"
}

func (mediaType MediaType) xmlName() xml.Name {
	return xml.Name{Local: "xml:" + mediaType.Name()}
}

func (mediaType MediaType) isSet() bool {
	return string(mediaType) != ""
}

func (mediaType MediaType) toAttr() xml.Attr {
	return xml.Attr{Name: mediaType.xmlName(), Value: mediaType.String()}
}

type Properties string

func (properties Properties) String() string {
	return string(properties)
}

func (Properties) Name() string {
	return "properties"
}

func (properties Properties) xmlName() xml.Name {
	return xml.Name{Local: "xml:" + properties.Name()}
}

func (properties Properties) isSet() bool {
	return string(properties) != ""
}

func (properties Properties) toAttr() xml.Attr {
	return xml.Attr{Name: properties.xmlName(), Value: properties.String()}
}
