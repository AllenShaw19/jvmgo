package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 读取dir目录下 .class文件
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		l
		return nil
	}
	return &DirEntry{absDir: absDir}
}

func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(e.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, e, err
}

func (e *DirEntry) String() string {
	return e.absDir
}
