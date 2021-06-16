package packagefile_test

import (
	"archive/zip"
	"testing"
	"time"

	. "github.com/quillpub/epub/package"
)

func TestChildrensLiterature(t *testing.T) {
	zip, err := zip.OpenReader("../testdata/childrens-literature.epub")
	if err != nil {
		t.Fatal(err)
	}

	pkg, err := GetPackage(zip, "EPUB/package.opf")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("check package", func(t *testing.T) {
		expectAttribute(t, Version("3.0"), pkg.Version)
		expectAttribute(t, UniqueIdentifier("id"), pkg.UniqueIdentifier)
		expectAttribute(t, TextDirection(""), pkg.TextDirection)
		expectAttribute(t, XMLLang(""), pkg.XMLLang)
		expectAttribute(t, ID(""), pkg.ID)
		expectAttribute(t, Prefix(""), pkg.Prefix)
	})
	t.Run("check metadata", func(t *testing.T) {
		metadata := pkg.Metadata
		expectCount(t, 1, len(metadata.Identifiers), "identifiers")
		identifier := metadata.Identifiers[0]
		expectAttribute(t, ID("id"), identifier.ID)
		expectString(t, "http://www.gutenberg.org/ebooks/25545", identifier.Text)

		expectCount(t, 2, len(metadata.Titles), "titles")
		title1 := metadata.Titles[0]
		expectAttribute(t, ID("t1"), title1.ID)
		expectString(t, "Children's Literature", title1.Text)
		title2 := metadata.Titles[1]
		expectAttribute(t, ID("t2"), title2.ID)
		expectString(t, "A Textbook of Sources for Teachers and Teacher-Training Classes", title2.Text)

		expectCount(t, 2, len(metadata.Creators), "creators")
		creator1 := metadata.Creators[0]
		expectAttribute(t, ID("curry"), creator1.ID)
		expectString(t, "Charles Madison Curry", creator1.Text)
		creator2 := metadata.Creators[1]
		expectAttribute(t, ID("clippinger"), creator2.ID)
		expectString(t, "Erle Elsworth Clippinger", creator2.Text)

		expectCount(t, 1, len(metadata.Languages), "languages")
		expectString(t, "en", metadata.Languages[0].Text)

		expectDate(t, time.Date(2008, 05, 20, 0, 0, 0, 0, time.UTC), metadata.Date)

		expectCount(t, 2, len(metadata.Subjects), "subjects")
		expectString(t, "Children -- Books and reading", metadata.Subjects[0].Text)
		expectString(t, "Children's literature -- Study and teaching", metadata.Subjects[1].Text)

		expectCount(t, 1, len(metadata.Sources), "sources")
		expectString(t, "http://www.gutenberg.org/files/25545/25545-h/25545-h.htm", metadata.Sources[0].Text)

		expectCount(t, 1, len(metadata.Rights), "rights")
		expectString(t, "Public domain in the USA.", metadata.Rights[0].Text)

		expectedMetas := []Meta{
			{Property: "dcterms:modified", Text: "2010-02-17T04:39:13Z"},
			{Refines: "#t1", Property: "title-type", Text: "main"},
			{Refines: "#t1", Property: "display-seq", Text: "1"},
			{Refines: "#t2", Property: "title-type", Text: "subtitle"},
			{Refines: "#t2", Property: "display-seq", Text: "2"},
			{Refines: "#curry", Property: "file-as", Text: "Curry, Charles Madison"},
			{Refines: "#clippinger", Property: "file-as", Text: "Clippinger, Erle Elsworth"},
		}
		expectCount(t, len(expectedMetas), len(metadata.Metas), "metas")
		for i, meta := range expectedMetas {
			expectAttribute(t, meta.Refines, metadata.Metas[i].Refines)
			expectAttribute(t, meta.Property, metadata.Metas[i].Property)
			expectString(t, meta.Text, metadata.Metas[i].Text)
		}
	})
}
