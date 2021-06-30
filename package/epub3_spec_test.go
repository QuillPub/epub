package packagefile_test

import (
	"archive/zip"
	"testing"

	. "github.com/quillpub/epub/package"
)

func TestEpub3Spec(t *testing.T) {
	zip, err := zip.OpenReader("../testdata/epub30-spec.epub")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := zip.Close()
		if err != nil {
			t.Errorf("close container file error: %s", err)
		}
	}()

	pkg, err := GetPackage(&zip.Reader, "EPUB/package.opf")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("check package", func(t *testing.T) {
		expectAttribute(t, Version("3.0"), pkg.Version)
		expectAttribute(t, UniqueIdentifier("uid"), pkg.UniqueIdentifier)
		expectAttribute(t, TextDirection(""), pkg.TextDirection)
		expectAttribute(t, XMLLang(""), pkg.XMLLang)
		expectAttribute(t, ID(""), pkg.ID)
		expectAttribute(t, Prefix(""), pkg.Prefix)
	})
	t.Run("check metadata", func(t *testing.T) {
		metadata := pkg.Metadata
		expectCount(t, 1, len(metadata.Identifiers), "identifiers")
		identifier := metadata.Identifiers[0]
		expectAttribute(t, ID("uid"), identifier.ID)
		expectString(t, "code.google.com.epub-samples.epub30-spec", identifier.Text)

		expectCount(t, 1, len(metadata.Titles), "titles")
		expectString(t, "EPUB 3.0 Specification", metadata.Titles[0].Text)

		expectCount(t, 1, len(metadata.Languages), "languages")
		expectString(t, "en", metadata.Languages[0].Text)

		expectCount(t, 1, len(metadata.Metas), "metas")
		modifiedMeta := metadata.Metas[0]
		expectAttribute(t, Property("dcterms:modified"), modifiedMeta.Property)
		expectString(t, "2012-02-27T16:38:35Z", modifiedMeta.Text)

		expectCount(t, 1, len(metadata.Creators), "creators")
		expectString(t, "EPUB 3 Working Group", metadata.Creators[0].Text)
	})
}
