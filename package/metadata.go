package packagefile

import (
	"encoding/xml"
	"time"
)

type Metadata struct {
	XMLName     xml.Name
	Identifiers []Identifier `xml:"identifier"`
	Titles      []Title      `xml:"title"`
	Languages   []Language   `xml:"language"`
	Metas       []Meta       `xml:"meta"`

	//Optional
	Creators     []Creator
	Descriptions []Description
	Publishers   []Publisher
	Contributors []Contributor
	Coverages    []Coverage
	Relations    []Relation
	Rights       []Rights
	Subjects     []Subject
	Types        []Type
	Links        []Link
	Sources      []Source
	Date         Date `xml:"date,omitempty"`
}

func (metadata Metadata) elements() []interface {
} {
	return []interface{}{
		&metadata.Identifiers,
		&metadata.Titles,
		&metadata.Languages,
		&metadata.Metas,
		&metadata.Creators,
		&metadata.Descriptions,
		&metadata.Publishers,
		&metadata.Contributors,
		&metadata.Coverages,
		&metadata.Relations,
		&metadata.Rights,
		&metadata.Subjects,
		&metadata.Types,
		&metadata.Sources,
		&metadata.Links,
		&metadata.Date,
	}
}

func (metadata *Metadata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	metadata.XMLName = start.Name

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "identifier":
				var identifier Identifier
				if err = d.DecodeElement(&identifier, &el); err != nil {
					return err
				}
				metadata.Identifiers = append(metadata.Identifiers, identifier)

			case "title":
				var title Title
				if err = d.DecodeElement(&title, &el); err != nil {
					return err
				}
				metadata.Titles = append(metadata.Titles, title)

			case "language":
				var language Language
				if err = d.DecodeElement(&language, &el); err != nil {
					return err
				}
				metadata.Languages = append(metadata.Languages, language)

			case "meta":
				var meta Meta
				if err = d.DecodeElement(&meta, &el); err != nil {
					return err
				}
				metadata.Metas = append(metadata.Metas, meta)

			case "creator":
				creator := Creator{DCMESOptionalElement{name: "creator"}}
				if err = d.DecodeElement(&creator, &el); err != nil {
					return err
				}
				metadata.Creators = append(metadata.Creators, creator)

			case "publisher":
				publisher := Publisher{DCMESOptionalElement{name: "publisher"}}
				if err = d.DecodeElement(&publisher, &el); err != nil {
					return err
				}
				metadata.Publishers = append(metadata.Publishers, publisher)

			case "description":
				description := Description{DCMESOptionalElement{name: "description"}}
				if err = d.DecodeElement(&description, &el); err != nil {
					return err
				}
				metadata.Descriptions = append(metadata.Descriptions, description)

			case "contributor":
				contributor := Contributor{DCMESOptionalElement{name: "contributor"}}
				if err = d.DecodeElement(&contributor, &el); err != nil {
					return err
				}
				metadata.Contributors = append(metadata.Contributors, contributor)

			case "coverage":
				coverage := Coverage{DCMESOptionalElement{name: "coverage"}}
				if err = d.DecodeElement(&coverage, &el); err != nil {
					return err
				}
				metadata.Coverages = append(metadata.Coverages, coverage)

			case "relation":
				relation := Relation{DCMESOptionalElement{name: "relation"}}
				if err = d.DecodeElement(&relation, &el); err != nil {
					return err
				}
				metadata.Relations = append(metadata.Relations, relation)

			case "rights":
				rights := Rights{DCMESOptionalElement{name: "rights"}}
				if err = d.DecodeElement(&rights, &el); err != nil {
					return err
				}
				metadata.Rights = append(metadata.Rights, rights)

			case "subject":
				subject := Subject{DCMESOptionalElement{name: "subject"}}
				if err = d.DecodeElement(&subject, &el); err != nil {
					return err
				}
				metadata.Subjects = append(metadata.Subjects, subject)

			case "type":
				t := Type{DCMESOptionalElement{name: "type"}}
				if err = d.DecodeElement(&t, &el); err != nil {
					return err
				}
				metadata.Types = append(metadata.Types, t)

			case "source":
				source := Source{DCMESOptionalElement{name: "source"}}
				if err = d.DecodeElement(&source, &el); err != nil {
					return err
				}
				metadata.Sources = append(metadata.Sources, source)

			case "link":
				var link Link
				if err = d.DecodeElement(&link, &el); err != nil {
					return err
				}
				metadata.Links = append(metadata.Links, link)

			case "date":
				if err = d.DecodeElement(&metadata.Date, &el); err != nil {
					return err
				}
			}
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

func (metadata Metadata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = metadata.XMLName
	start.Name.Space = ""
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns:dc"}, Value: dcNamespace},
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	for _, element := range metadata.elements() {
		err = e.Encode(element)
		if err != nil {
			return err
		}
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}

const ISO8601Date = `2006-01-02`

type Date struct {
	T *time.Time
}

func (date *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.CharData:
			text := string(el)
			t, err := time.Parse(ISO8601Date, text)
			if err != nil {
				return err
			}
			date.T = &t
		case xml.StartElement:
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

func (date *Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if date.T == nil {
		return nil
	}
	start.Name.Local = "dc:date"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.CharData(date.T.Format(ISO8601Date)))
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}
	return nil
}
