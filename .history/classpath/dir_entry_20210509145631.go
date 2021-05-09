package classpath

import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) (*DirEntry, error) {
	absDir, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &DirEntry{absDir: absDir}, nil
}

func (e *DirEntry) readClass(className string) 