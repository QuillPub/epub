package container_test

import (
	"archive/zip"
	"testing"

	. "github.com/quillpub/epub/container"
)

func TestGetContentsPath(t *testing.T) {
	testCases := []struct {
		filePath     string
		expectedPath string
	}{
		{filePath: "../testdata/epub30-spec.epub", expectedPath: "EPUB/package.opf"},
		{filePath: "../testdata/accessible_epub_3.epub", expectedPath: "EPUB/package.opf"},
		{filePath: "../testdata/haruko-html-jpeg.epub", expectedPath: "OPS/package.opf"},
		{filePath: "../testdata/georgia-pls-ssml.epub", expectedPath: "EPUB/package.opf"},
		{filePath: "../testdata/figure-gallery-bindings.epub", expectedPath: "EPUB/package.opf"},
		{filePath: "../testdata/cc-shared-culture.epub", expectedPath: "EPUB/package.opf"},
		{filePath: "../testdata/childrens-literature.epub", expectedPath: "EPUB/package.opf"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.filePath, func(t *testing.T) {
			zip, err := zip.OpenReader(testCase.filePath)
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				err := zip.Close()
				if err != nil {
					t.Errorf("close container file error: %s", err)
				}
			}()

			path, err := GetContentsPath(zip)
			if err != nil {
				t.Error(err)
			}
			if path != testCase.expectedPath {
				t.Errorf("expected path %s, got %s", testCase.expectedPath, path)
			}
		})
	}
}
