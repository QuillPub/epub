package packagefile_test

import (
	"archive/zip"
	"testing"

	. "github.com/quillpub/epub/package"
)

func TestGeorgia(t *testing.T) {
	zip, err := zip.OpenReader("../testdata/georgia-pls-ssml.epub")
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
		expectAttribute(t, XMLLang("en-US"), pkg.XMLLang)
		expectAttribute(t, ID(""), pkg.ID)
		expectAttribute(t, Prefix(""), pkg.Prefix)
	})
	t.Run("check metadata", func(t *testing.T) {
		metadata := pkg.Metadata
		expectCount(t, 1, len(metadata.Identifiers), "identifiers")
		identifier := metadata.Identifiers[0]
		expectAttribute(t, ID("uid"), identifier.ID)
		expectString(t, "code.google.com.epub-samples.georgia-pls-ssml", identifier.Text)

		expectedTitles := []Title{
			{ID: "t1", Text: "Georgia"},
			{ID: "full-title", Text: "Encyclopaedia Britannica, 11th Edition, Volume 11, Slice 7 / Georgia"},
			{ID: "t2", Text: "Encyclopaedia Britannica"},
			{ID: "t3", Text: "11th Edition"},
		}
		expectCount(t, len(expectedTitles), len(metadata.Titles), "titles")
		for i, expectedTitle := range expectedTitles {
			expectAttribute(t, expectedTitle.ID, metadata.Titles[i].ID)
			expectString(t, expectedTitle.Text, metadata.Titles[i].Text)
		}

		expectCount(t, 1, len(metadata.Creators), "creators")
		creator1 := metadata.Creators[0]
		expectAttribute(t, ID("aut1"), creator1.ID)
		expectString(t, "Various", creator1.Text)

		expectCount(t, 1, len(metadata.Languages), "languages")
		expectString(t, "en-US", metadata.Languages[0].Text)

		expectedMetas := []Meta{
			{Refines: "#t1", Property: "title-type", Text: "main"},
			{Refines: "#t1", Property: "display-seq", Text: "3"},
			{Refines: "#full-title", Property: "title-type", Text: "expanded"},
			{Refines: "#t2", Property: "title-type", Text: "collection"},
			{Refines: "#t2", Property: "display-seq", Text: "1"},
			{Refines: "#t3", Property: "title-type", Text: "edition"},
			{Refines: "#t3", Property: "display-seq", Text: "2"},
			{Property: "dcterms:modified", ID: "mod", Text: "2012-02-07T16:38:35Z"},
			{Refines: "#aut1", Property: "role", Scheme: "marc:relators", Text: "aut"},
			{Property: "dcterms:source", Text: "http://www.gutenberg.org/files/37523/37523-h/37523-h.htm"},
		}
		expectCount(t, len(expectedMetas), len(metadata.Metas), "metas")
		for i, meta := range expectedMetas {
			expectAttribute(t, meta.Refines, metadata.Metas[i].Refines)
			expectAttribute(t, meta.Property, metadata.Metas[i].Property)
			expectAttribute(t, meta.ID, metadata.Metas[i].ID)
			expectAttribute(t, meta.Scheme, metadata.Metas[i].Scheme)
			expectString(t, meta.Text, metadata.Metas[i].Text)
		}
	})
}
