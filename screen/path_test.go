package screen

import (
	"fmt"
	"grover/files"
	"strings"
	"testing"
)

func TestLoadExisting(t *testing.T) {
	files.Mkdir("data")
	defer files.Rmdir("data")

	list := LoadPaths("main.go")
	fmt.Println(1, GatherPaths(list))
	if len(list) != 1 {
		t.Fatal("wrong length")
	}
	WritePaths(list)

	// -----

	list = LoadPaths("main2.go")
	fmt.Println(2, GatherPaths(list))
	if len(list) != 2 {
		t.Fatal("wrong length")
	}
	if list[0].Fullpath != "main2.go" {
		t.Fatal("wrong order")
	}
	WritePaths(list)

	// -----

	list = LoadPaths("main3.go")
	fmt.Println(3, GatherPaths(list))
	if len(list) != 3 {
		t.Fatal("wrong length")
	}
	if list[0].Fullpath != "main3.go" {
		t.Fatal("wrong order")
	}
	WritePaths(list)

	// -----

	list = LoadPaths("main2.go")
	fmt.Println(4, GatherPaths(list))
	if len(list) != 3 {
		t.Fatal("wrong length")
	}
	if list[0].Fullpath != "main2.go" {
		t.Fatal("wrong order")
	}
	WritePaths(list)

}

func GatherPaths(list []*Path) string {
	buff := []string{}
	for _, p := range list {
		buff = append(buff, p.Fullpath)
	}
	return strings.Join(buff, ",")
}
