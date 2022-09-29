package screen

import (
	"grover/files"
	"strings"
)

type Path struct {
	Fullpath string
	Filename string
}

func NewPath(s string) *Path {
	tokens := strings.Split(s, "/")

	p := Path{}
	p.Fullpath = s
	p.Filename = tokens[len(tokens)-1]
	return &p
}

func LoadPaths() []*Path {
	s := files.ReadFile("data/paths.txt")
	lines := strings.Split(s, "\n")
	paths := []*Path{}
	for _, line := range lines {
		paths = append(paths, NewPath(line))
	}
	return paths
}

func WritePaths(paths []*Path) {
	buff := []string{}
	for _, p := range paths {
		buff = append(buff, p.Fullpath)
	}
	payload := strings.Join(buff, "\n")
	files.SaveFile("data/paths.txt", payload)
}
