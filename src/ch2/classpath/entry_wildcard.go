package classpath

import (
	"os"
	"path/filepath"
	"strings"
	"fmt"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path) - 1] // remote *
	compopositeEntry := []Entry{}

	fmt.Printf("new Wildcard path %s, baseDir %s\n", path, baseDir)
	walkFn := func(path string, info os.FileInfo, err error) error {
		//fmt.Printf("path %s, \n", path)
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compopositeEntry = append(compopositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)

	fmt.Println(len(compopositeEntry))
	return compopositeEntry
}
