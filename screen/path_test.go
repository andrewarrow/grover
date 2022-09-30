package screen

import (
	"grover/files"
	"testing"
)

func TestLoadExisting(t *testing.T) {
	files.Mkdir("data")
	LoadPaths("main.go")
	files.Rmdir("data")
}
