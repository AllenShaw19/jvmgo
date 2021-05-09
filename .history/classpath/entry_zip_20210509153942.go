package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) (*ZipEntry, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &ZipEntry{absPath: absPath}, nil
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(e.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
		}
	}
}