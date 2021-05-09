package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// WildcardEntry实际上是CompositeEntry
type WildcardEntry CompositeEntry

func newWildcardEntry(path string) (WildcardEntry, error) {
	baseDir := path[:len(path)-1]
	wildcardEntry := []Entry{}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry, err := newZipEntry(path)
			if err != nil {
				return err
			}
			wildcardEntry = append(wildcardEntry, jarEntry)
		}
		return nil
	}

	filepath.
}
