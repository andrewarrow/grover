package main

import (
	"SwiftExplorer/swift"
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
		if strings.HasSuffix(path, ".swift") {
			fmt.Println(path)
			swift.ReadFile(path)
		}
		return nil
	}

	filepath.Walk(dir, f)
}
