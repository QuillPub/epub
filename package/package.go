package packagefile

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"log"
)

// Package contains all the information about the epub
// See https://www.w3.org/publishing/epub3/epub-packages.html#sec-package-elem
type Package struct {
	XMLName xml.Name `xml:"package"`
	// Required
	Version          Version
	UniqueIdentifier UniqueIdentifier

	// Optional
	TextDirection TextDirection
	ID            ID
	Prefix        Prefix
	XMLLang       XMLLang

	Metadata Metadata
	//Manifest []Item   `xml:"manifest"`
	//Spine    []string `xml:"spine`
	//Guide
	//Collection
}

func (pkg Package) optionalAttributes() []optionalAttribute {
	return []optionalAttribute{
		pkg.TextDirection,
		pkg.ID,
		pkg.Prefix,
		pkg.XMLLang,
	}
}

func (pkg *Package) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	pkg.XMLName = start.Name
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case pkg.Version.Name():
			pkg.Version = Version(attr.Value)
		case pkg.UniqueIdentifier.Name():
			pkg.UniqueIdentifier = UniqueIdentifier(attr.Value)
		case pkg.TextDirection.Name():
			pkg.TextDirection = TextDirection(attr.Value)
		case pkg.ID.Name():
			pkg.ID = ID(attr.Value)
		case pkg.Prefix.Name():
			pkg.Prefix = Prefix(attr.Value)
		case pkg.XMLLang.Name():
			if attr.Name.Space != xmlNamespace {
				continue
			}
			pkg.XMLLang = XMLLang(attr.Value)
		}
	}

	for {
		token, err := d.Token()
		if err != nil {
			return err
		}

		switch el := token.(type) {
		case xml.StartElement:
			if el.Name.Local == "metadata" {
				var metadata Metadata
				if err = d.DecodeElement(&metadata, &el); err != nil {
					return err
				}
				pkg.Metadata = metadata
			}
		case xml.EndElement:
			if el == start.End() {
				return nil
			}
		}
	}
}

func (pkg Package) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = pkg.XMLName
	start.Attr = []xml.Attr{
		pkg.Version.toAttr(),
		pkg.UniqueIdentifier.toAttr(),
	}
	for _, attr := range pkg.optionalAttributes() {
		if attr.isSet() {
			start.Attr = append(start.Attr, attr.toAttr())
		}
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	err = e.Encode(pkg.Metadata)
	if err != nil {
		return err
	}

	err = e.EncodeToken(xml.EndElement{Name: start.Name})
	if err != nil {
		return err
	}

	return nil
}

// GetPackage retrieves the package from an epub file (which is a zip) and
// parses it using ParsePackage
func GetPackage(zipFile *zip.Reader, path string) (*Package, error) {
	file, err := openPackage(zipFile, path)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("close container file error: %s", err)
		}
	}()
	return ParsePackage(file)
}

func openPackage(zip *zip.Reader, path string) (io.ReadCloser, error) {
	return zip.Open(path)
}

// ParsePackage parses a package file into a Package.
func ParsePackage(file io.Reader) (*Package, error) {
	var p Package
	decoder := xml.NewDecoder(file)
	err := decoder.Decode(&p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Version contains the epub standard version
type Version string

func (version Version) String() string {
	return string(version)
}

// Name gives the xml attribute name
func (Version) Name() string {
	return "version"
}

func (version Version) xmlName() xml.Name {
	return xml.Name{Local: version.Name()}
}

func (version Version) isSet() bool {
	return string(version) != ""
}

func (version Version) toAttr() xml.Attr {
	return xml.Attr{Name: version.xmlName(), Value: version.String()}
}

// UniqueIdentifier idenitifies the primary Identifier in the metadata
type UniqueIdentifier string

func (uid UniqueIdentifier) String() string {
	return string(uid)
}

// Name gives the xml attribute name
func (UniqueIdentifier) Name() string {
	return "unique-identifier"
}

func (uid UniqueIdentifier) xmlName() xml.Name {
	return xml.Name{Local: uid.Name()}
}

func (uid UniqueIdentifier) isSet() bool {
	return string(uid) != ""
}

func (uid UniqueIdentifier) toAttr() xml.Attr {
	return xml.Attr{Name: uid.xmlName(), Value: uid.String()}
}

// See https://www.w3.org/publishing/epub3/epub-packages.html#sec-prefix-attr
type Prefix string

func (prefix Prefix) String() string {
	return string(prefix)
}

// Name gives the xml attribute name
func (Prefix) Name() string {
	return "prefix"
}

func (prefix Prefix) xmlName() xml.Name {
	return xml.Name{Local: prefix.Name()}
}

func (prefix Prefix) isSet() bool {
	return string(prefix) != ""
}

func (prefix Prefix) toAttr() xml.Attr {
	return xml.Attr{Name: prefix.xmlName(), Value: prefix.String()}
}
