package main

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
		if strings.Index(path, "/.git") == -1 {
			fmt.Println(path, info.Size())
		}
		return nil
	}

	filepath.Walk(dir, f)
}
