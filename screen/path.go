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

func LoadPaths(fullpath string) ([]*Path, bool) {
	s := files.ReadFile("data/paths.txt")
	lines := strings.Split(s, "\n")
	paths := []*Path{}
	mapForDups := map[string]bool{}
	for _, line := range lines {
		paths = append(paths, NewPath(line))
		mapForDups[line] = true
	}
	return paths, mapForDups[fullpath]
}

func WritePaths(paths []*Path) {
	buff := []string{}
	mapForDups := map[string]bool{}
	for _, p := range paths {
		if mapForDups[p.Fullpath] {
			continue
		}
		buff = append(buff, p.Fullpath)
		mapForDups[p.Fullpath] = true
	}
	payload := strings.Join(buff, "\n")
	files.SaveFile("data/paths.txt", payload)
}
