package screen

import (
	"grover/files"
	"strings"
	"testing"
)

func TestLoadExisting(t *testing.T) {
	files.Mkdir("data")
	defer files.Rmdir("data")

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

	list = LoadPaths("main3.go")
	if list[0].Fullpath != "main3.go" {
		t.Fatal("wrong order")
	}
	WritePaths(list)

	list = LoadPaths("main2.go")
	if list[0].Fullpath != "main2.go" {
		t.Fatal("wrong order", PrintPaths(list))
	}
	WritePaths(list)

}

func PrintPaths(list []*Path) string {
	buff := []string{}
	for _, p := range list {
		buff = append(buff, p.Fullpath)
	}
	return strings.Join(buff, ",")
}
