package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 读取解析zip,jar文件形式的类
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry{
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &ZipEntry{absPath: absPath}
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// TODO: 待优化，无需每次都zipOpenReader,是否可以缓存？
	r, err := zip.OpenReader(e.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, e, err
		}
	}

	return nil, nil, errors.New("class not found " + className)
}

func (e *ZipEntry) String() string {
	return e.absPath
}
