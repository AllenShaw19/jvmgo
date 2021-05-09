package classpath

import "os"

// WildcardEntry实际上是CompositeEntry
type WildcardEntry CompositeEntry

func newWildcardEntry(path string) WildcardEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir( path != baseDir {}

	}
}
