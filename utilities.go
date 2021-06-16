package epub

import "archive/zip"

func openFile(path string) (*zip.ReadCloser, error) {
	return zip.OpenReader(path)
}
