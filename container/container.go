package container

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"io"
	"log"
)

type Rootfile struct {
	Path      string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}
type Container struct {
	XMLName    xml.Name   `xml:"container"`
	Roootfiles []Rootfile `xml:"rootfiles>rootfile"`
}

const MediaTypePackageFile = `application/oebps-package+xml`

var ErrContentsPathNotFound = errors.New("contents path not found")

func GetContentsPath(zip *zip.Reader) (string, error) {
	file, err := zip.Open("META-INF/container.xml")
	if err != nil {
		return "", err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("close container file error: %s", err)
		}
	}()
	container, err := parseContainer(file)
	if err != nil {
		return "", err
	}
	return extractContentsPath(container)
}

func parseContainer(r io.Reader) (*Container, error) {
	var container Container
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&container)
	if err != nil {
		return nil, err
	}
	return &container, nil
}

func extractContentsPath(container *Container) (string, error) {
	var contentsPath string
	for _, rootfile := range container.Roootfiles {
		if rootfile.MediaType == MediaTypePackageFile {
			contentsPath = rootfile.Path
		}
	}
	if contentsPath == "" {
		return "", ErrContentsPathNotFound
	}
	return contentsPath, nil
}
