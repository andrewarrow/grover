package screen

import "strings"

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
