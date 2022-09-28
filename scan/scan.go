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
	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".kt") {
			lower := strings.ToLower(path)
			if strings.Contains(lower, lowerFilter) {
				fmt.Println(path)
			}
		}
		return nil
	}

	filepath.Walk(dir, f)
	screen.Filter()
}

func OneFile(path string) {
	//characters := swift.ReadFile(path)
	//swift.ProcessCharacters(characters)
}
