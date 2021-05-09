package classpath

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jvmgo/logger"
)

// WildcardEntry实际上是CompositeEntry
type WildcardEntry CompositeEntry

func newWildcardEntry(path string) CompositeEntry {
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
			jarEntry := newZipEntry(path)
			wildcardEntry = append(wildcardEntry, jarEntry)
		}
		return nil
	}

	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		logger.Error("walk dir err %v", err)
	}
	return wildcardEntry
}
