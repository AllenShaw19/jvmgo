package classpath

import (
	"errors"
	"strings"
)

//
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (e CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// TODO: 这个实现有点不太优雅，能否直接找到class
	for _, entry := range e {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found " + className)
}

func ()