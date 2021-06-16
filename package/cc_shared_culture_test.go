package packagefile_test

import (
	"archive/zip"
	"testing"

	. "github.com/quillpub/epub/package"
)

func TestCCSharedCulture(t *testing.T) {
	zip, err := zip.OpenReader("../testdata/cc-shared-culture.epub")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := zip.Close()
		if err != nil {
			t.Errorf("close container file error: %s", err)
		}
	}()

	pkg, err := GetPackage(zip, "EPUB/package.opf")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("check package", func(t *testing.T) {
		expectAttribute(t, Version("3.0"), pkg.Version)
		expectAttribute(t, UniqueIdentifier("uid"), pkg.UniqueIdentifier)
		expectAttribute(t, TextDirection(""), pkg.TextDirection)
		expectAttribute(t, XMLLang("en"), pkg.XMLLang)
		expectAttribute(t, ID(""), pkg.ID)
		expectAttribute(t, Prefix("cc: http://creativecommons.org/ns#"), pkg.Prefix)
	})
	t.Run("check metadata", func(t *testing.T) {
		metadata := pkg.Metadata
		expectCount(t, 1, len(metadata.Identifiers), "identifiers")
		identifier := metadata.Identifiers[0]
		expectAttribute(t, ID("uid"), identifier.ID)
		expectString(t, "code.google.com.epub-samples.cc-shared-culture", identifier.Text)

		expectCount(t, 1, len(metadata.Titles), "titles")
		title := metadata.Titles[0]
		expectAttribute(t, ID("title"), title.ID)
		expectString(t, "Creative Commons - A Shared Culture", title.Text)

		expectCount(t, 1, len(metadata.Languages), "languages")
		expectString(t, "en-US", metadata.Languages[0].Text)

		expectCount(t, 2, len(metadata.Metas), "metas")
		modifiedMeta := metadata.Metas[0]
		expectAttribute(t, Property("dcterms:modified"), modifiedMeta.Property)
		expectString(t, "2012-01-20T12:47:00Z", modifiedMeta.Text)
		attributionMeta := metadata.Metas[1]
		expectAttribute(t, Property("cc:attributionURL"), attributionMeta.Property)
		expectString(t, "http://creativecommons.org/videos/a-shared-culture", attributionMeta.Text)

		expectCount(t, 1, len(metadata.Creators), "creators")
		expectString(t, "Jesse Dylan", metadata.Creators[0].Text)

		expectCount(t, 1, len(metadata.Publishers), "publishers")
		expectString(t, "Creative Commons", metadata.Publishers[0].Text)

		expectCount(t, 1, len(metadata.Contributors), "contributors")
		expectString(t, "mgylling", metadata.Contributors[0].Text)

		expectCount(t, 1, len(metadata.Descriptions), "descriptions")
		expectString(t, "Multiple video tests (see Navigation Document (toc) for details)", metadata.Descriptions[0].Text)

		expectCount(t, 1, len(metadata.Rights), "rights")
		expectString(t, "This work is licensed under a Creative Commons Attribution-Noncommercial-Share Alike (CC BY-NC-SA) license.", metadata.Rights[0].Text)

		expectCount(t, 3, len(metadata.Links), "links")
		link1 := metadata.Links[0]
		expectAttribute(t, Rel("cc:license"), link1.Rel)
		expectAttribute(t, Href("http://creativecommons.org/licenses/by-nc-sa/3.0/"), link1.Href)
		link2 := metadata.Links[1]
		expectAttribute(t, Refines("#img1"), link2.Refines)
		expectAttribute(t, Rel("cc:license"), link2.Rel)
		expectAttribute(t, Href("http://creativecommons.org/licenses/by-nc/2.0/"), link2.Href)
		link3 := metadata.Links[2]
		expectAttribute(t, Refines("#img2"), link3.Refines)
		expectAttribute(t, Rel("cc:license"), link3.Rel)
		expectAttribute(t, Href("http://creativecommons.org/licenses/by-nc-sa/2.0/"), link3.Href)
	})
}
