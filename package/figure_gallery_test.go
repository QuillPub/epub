package packagefile_test

import (
	"archive/zip"
	"testing"
	"time"

	. "github.com/quillpub/epub/package"
)

func TestFigureGallery(t *testing.T) {
	zip, err := zip.OpenReader("../testdata/figure-gallery-bindings.epub")
	if err != nil {
		t.Fatal(err)
	}

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
		expectAttribute(t, Prefix("cc: http://creativecommons.org/ns#"), pkg.Prefix)
	})
	t.Run("check metadata", func(t *testing.T) {
		metadata := pkg.Metadata

		expectCount(t, 1, len(metadata.Identifiers), "identifiers")
		identifier := metadata.Identifiers[0]
		expectAttribute(t, ID("uid"), identifier.ID)
		expectString(t, "code.google.com.epub-samples.figure-gallery", identifier.Text)

		expectCount(t, 1, len(metadata.Titles), "titles")
		title := metadata.Titles[0]
		expectString(t, "EPUB Figure Gallery", title.Text)

		expectCount(t, 1, len(metadata.Creators), "creators")
		expectString(t, "E.L. Starr", metadata.Creators[0].Text)

		expectCount(t, 1, len(metadata.Languages), "languages")
		expectString(t, "en-US", metadata.Languages[0].Text)

		expectDate(t, time.Date(2012, 01, 17, 0, 0, 0, 0, time.UTC), pkg.Metadata.Date)

		expectCount(t, 1, len(metadata.Rights), "rights")
		expectString(t, "This work is shared with the public using the Attribution-ShareAlike 3.0 Unported (CC BY-SA 3.0) license.", metadata.Rights[0].Text)

		expectCount(t, 3, len(metadata.Metas), "metas")
		modifiedMeta := metadata.Metas[0]
		expectAttribute(t, Property("dcterms:modified"), modifiedMeta.Property)
		expectString(t, "2012-04-05T21:03:00Z", modifiedMeta.Text)

		attributionMeta := metadata.Metas[1]
		expectAttribute(t, Property("cc:attributionURL"), attributionMeta.Property)
		expectString(t, "http://code.google.com/p/epub-samples/", attributionMeta.Text)
		// meta[2] is there for epub2 compatibility, which we do not currently support. Ignore it for now

		expectedLinks := []Link{
			{Rel: "cc:license", Href: "http://creativecommons.org/licenses/by-sa/3.0/"},
			{Rel: "cc:attributionURL", Refines: "#cover", Href: "http://www.flickr.com/photos/smithsonian/2941526052"},
			{Rel: "cc:attributionURL", Refines: "#fig1", Href: "http://www.flickr.com/photos/gsfc/5837033098"},
			{Rel: "cc:attributionURL", Refines: "#fig2", Href: "http://www.flickr.com/photos/gsfc/5836481331"},
			{Rel: "cc:attributionURL", Refines: "#fig3", Href: "http://www.flickr.com/photos/gsfc/5837030890"},
			{Rel: "cc:attributionURL", Refines: "#fig4", Href: "http://www.flickr.com/photos/gsfc/5837031188"},
			{Rel: "cc:attributionURL", Refines: "#fig5", Href: "http://www.flickr.com/photos/gsfc/5836482263"},
			{Rel: "cc:attributionURL", Refines: "#fig6", Href: "http://www.flickr.com/photos/gsfc/5837032132"},
			{Rel: "cc:attributionURL", Refines: "#fig7", Href: "http://www.flickr.com/photos/gsfc/5837032522"},
			{Rel: "cc:attributionURL", Refines: "#fig8", Href: "http://www.flickr.com/photos/gsfc/5837032862"},
			{Rel: "cc:attributionURL", Refines: "text.xhtml#moon-text", Href: "http://en.wikipedia.org/wiki/Lunar_phase"},
			{Rel: "cc:attributionURL", Refines: "text.xhtml#moon-figures", Href: "http://www.flickr.com/photos/gsfc/sets/72157626962784062/"},
			{Rel: "cc:attributionURL", Refines: "#moon-phases-xml", Href: "http://www.flickr.com/photos/gsfc/sets/72157626962784062/"},
			{Rel: "cc:attributionURL", Refines: "text.xhtml#moon-figures", Href: "http://www.flickr.com/photos/gsfc/sets/72157626962784062/"},
		}
		expectCount(t, len(expectedLinks), len(metadata.Links), "links")
		for i, expectedLink := range expectedLinks {
			expectAttribute(t, expectedLink.Rel, metadata.Links[i].Rel)
			expectAttribute(t, expectedLink.Refines, metadata.Links[i].Refines)
			expectAttribute(t, expectedLink.Href, metadata.Links[i].Href)
		}
	})
}
