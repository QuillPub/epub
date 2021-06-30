package packagefile_test

import (
	"archive/zip"
	"testing"
	"time"

	. "github.com/quillpub/epub/package"
)

func TestAccessibleEpub3(t *testing.T) {
	zip, err := zip.OpenReader("../testdata/accessible_epub_3.epub")
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
		expectAttribute(t, UniqueIdentifier("pub-identifier"), pkg.UniqueIdentifier)
		expectAttribute(t, TextDirection(""), pkg.TextDirection)
		expectAttribute(t, XMLLang("en"), pkg.XMLLang)
		expectAttribute(t, ID(""), pkg.ID)
		expectAttribute(t, Prefix(""), pkg.Prefix)
	})
	t.Run("check metadata", func(t *testing.T) {
		metadata := pkg.Metadata
		expectCount(t, 1, len(metadata.Identifiers), "identifiers")
		pubIdentifier := metadata.Identifiers[0]
		expectAttribute(t, ID("pub-identifier"), pubIdentifier.ID)
		expectString(t, "urn:isbn:9781449328030", pubIdentifier.Text)

		expectCount(t, 1, len(metadata.Titles), "titles")
		pubTitle := metadata.Titles[0]
		expectAttribute(t, ID("pub-title"), pubTitle.ID)
		expectString(t, "Accessible EPUB 3", pubTitle.Text)

		expectCount(t, 1, len(metadata.Languages), "languages")
		pubLanguage := metadata.Languages[0]
		expectAttribute(t, ID("pub-language"), pubLanguage.ID)
		expectString(t, "en", pubLanguage.Text)

		expectCount(t, 1, len(metadata.Metas), "metas")
		modifiedMeta := metadata.Metas[0]
		expectAttribute(t, Property("dcterms:modified"), modifiedMeta.Property)
		expectString(t, "2012-10-24T15:30:00Z", modifiedMeta.Text)

		expectCount(t, 1, len(metadata.Creators), "creators")
		creator := metadata.Creators[0]
		expectAttribute(t, ID("pub-creator12"), creator.ID)
		expectString(t, "Matt Garrish", creator.Text)

		expectCount(t, 1, len(metadata.Publishers), "publishers")
		expectString(t, "O’Reilly Media, Inc.", metadata.Publishers[0].Text)

		expectCount(t, 6, len(metadata.Contributors), "contributors")
		expectString(t, "O’Reilly Production Services", metadata.Contributors[0].Text)
		expectString(t, "David Futato", metadata.Contributors[1].Text)
		expectString(t, "Robert Romano", metadata.Contributors[2].Text)
		expectString(t, "Brian Sawyer", metadata.Contributors[3].Text)
		expectString(t, "Dan Fauxsmith", metadata.Contributors[4].Text)
		expectString(t, "Karen Montgomery", metadata.Contributors[5].Text)

		expectCount(t, 1, len(metadata.Rights), "rights")
		expectString(t, "Copyright © 2012 O’Reilly Media, Inc", metadata.Rights[0].Text)

		expectDate(t, time.Date(2012, 02, 20, 0, 0, 0, 0, time.UTC), metadata.Date)

	})
}
