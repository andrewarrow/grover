package scan

import (
	"fmt"
	"grover/screen"
	"os"
	"path/filepath"
	"strings"
)

func Scan(dir, filter string) {

	lowerFilter := strings.ToLower(filter)
	paths := []*screen.Path{}

	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".kt") {
			lower := strings.ToLower(path)
			if strings.Contains(lower, lowerFilter) {
				fmt.Println(path)
				paths = append(paths, screen.NewPath(path))
			}
		}
		return nil
	}

	filepath.Walk(dir, f)
	screen.Filter(paths)
}

func OneFile(path string) {
	//characters := swift.ReadFile(path)
	//swift.ProcessCharacters(characters)
}
