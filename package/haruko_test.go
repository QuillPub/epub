package packagefile_test

import (
	"archive/zip"
	"testing"

	. "github.com/quillpub/epub/package"
)

func TestHaruko(t *testing.T) {
	zip, err := zip.OpenReader("../testdata/haruko-html-jpeg.epub")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := zip.Close()
		if err != nil {
			t.Errorf("close container file error: %s", err)
		}
	}()

	pkg, err := GetPackage(&zip.Reader, "OPS/package.opf")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("check package", func(t *testing.T) {
		expectAttribute(t, Version("3.0"), pkg.Version)
		expectAttribute(t, UniqueIdentifier("ootuya-id"), pkg.UniqueIdentifier)
		expectAttribute(t, TextDirection(""), pkg.TextDirection)
		expectAttribute(t, XMLLang("ja"), pkg.XMLLang)
		expectAttribute(t, ID(""), pkg.ID)
		expectAttribute(t, Prefix("rendition: http://www.idpf.org/vocab/rendition/#"), pkg.Prefix)
	})
	t.Run("check metadata", func(t *testing.T) {
		metadata := pkg.Metadata
		expectCount(t, 1, len(metadata.Identifiers), "identifiers")
		identifier := metadata.Identifiers[0]
		expectAttribute(t, ID("ootuya-id"), identifier.ID)
		expectString(t, "urn:uuid:A649F639-6C1F-1014-8CC3-F813564D7508", identifier.Text)

		expectCount(t, 1, len(metadata.Titles), "titles")
		expectString(t, "ハルコさんの彼氏", metadata.Titles[0].Text)

		expectCount(t, 1, len(metadata.Languages), "languages")
		expectString(t, "ja-jp", metadata.Languages[0].Text)

		expectedMetas := []Meta{
			{Property: "dcterms:modified", Text: "2012-05-24T00:00:00Z"},
			{Property: "rendition:layout", Text: "pre-paginated"},
			{Property: "rendition:orientation", Text: "portrait"},
			{Property: "rendition:spread", Text: "landscape"},
		}
		expectCount(t, len(expectedMetas), len(metadata.Metas), "metas")
		for i, meta := range expectedMetas {
			expectAttribute(t, meta.Property, metadata.Metas[i].Property)
			expectString(t, meta.Text, metadata.Metas[i].Text)
		}
	})
}
