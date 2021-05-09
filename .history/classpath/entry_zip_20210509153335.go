package classpath

import "path/filepath"

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) (*ZipEntry, error) {
	absPath, err := filepath.Abs(path)
	
}