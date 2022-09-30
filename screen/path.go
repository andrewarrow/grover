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

func LoadPaths(fullpath string) []*Path {
	s := files.ReadFile("data/paths.txt")
	path := NewPath(fullpath)
	if len(s) == 0 {
		return []*Path{path}
	}
	lines := strings.Split(s, "\n")
	paths := []*Path{}
	for _, line := range lines {
		if line == fullpath {
			continue
		}
		paths = append(paths, NewPath(line))
	}

	paths = append([]*Path{path}, paths...)
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
