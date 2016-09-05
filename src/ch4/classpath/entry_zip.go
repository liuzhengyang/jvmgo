package classpath

import (
	"path/filepath"
	"io/ioutil"
	"archive/zip"
	"fmt"
	"errors"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	fmt.Printf("abs path %s\n", self.absPath)
	r, err := zip.OpenReader(self.absPath)
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

			return data, self, nil
		}
	}
	return nil, nil, errors.New("error")
}

func (self *ZipEntry) String() string {
	return "ni" + self.absPath
}
