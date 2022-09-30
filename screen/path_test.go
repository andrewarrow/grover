package screen

import (
	"grover/files"
	"testing"
)

func TestLoadExisting(t *testing.T) {
	files.Mkdir("data")

	list := LoadPaths("main.go")
	if len(list) != 1 {
		t.Fatal("wrong length")
	}
	WritePaths(list)

	list = LoadPaths("main2.go")
	if len(list) != 2 {
		t.Fatal("wrong length")
	}
	if list[0].Fullpath != "main2.go" {
		t.Fatal("wrong order")
	}

	WritePaths(list)

	files.Rmdir("data")
}
