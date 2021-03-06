package classpath

import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry, {
	absDir, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
}