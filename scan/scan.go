package scan

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Scan(dir string) {

	f := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".kt") {
			fmt.Println(path)
		}
		return nil
	}

	filepath.Walk(dir, f)
}

func OneFile(path string) {
	//characters := swift.ReadFile(path)
	//swift.ProcessCharacters(characters)
}
