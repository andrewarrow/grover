package swift

import (
	"regexp"
)

type SwiftFile struct {
	Classes []Class
}

type Class struct {
	Name      string
	Functions []Function
}

type Function struct {
	Name string
}

func NewClass(name string) Class {
	c := Class{}
	m := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	c.Name = m.ReplaceAllString(name, "")
	return c
}
