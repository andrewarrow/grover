package swift

import (
	"regexp"
)

type Block struct {
	Text   string
	Blocks []*Block
}

type Class struct {
	Name      string
	Functions []Function
}

type Function struct {
	Name string
}

func NewBlock(text string) Block {
	b := Block{}
	b.Text = text
	return b
}

func NewClass(name string) Class {
	c := Class{}
	m := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
	c.Name = m.ReplaceAllString(name, "")
	return c
}
