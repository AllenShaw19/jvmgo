package classpath

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jvmgo/logger"
)

// WildcardEntry实际上是CompositeEntry
type WildcardEntry CompositeEntry

func newWildcardEntry(path string) WildcardEntry {
	baseDir := path[:len(path)-1]
	wildcardEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
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

	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		logger.Error("")
	}
	return wildcardEntry
}
