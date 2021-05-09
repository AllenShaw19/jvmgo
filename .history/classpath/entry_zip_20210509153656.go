package classpath

import (
	"archive/zip"
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
	r, err := zip.OpenReader
}