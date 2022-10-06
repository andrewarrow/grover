package swift

import (
	"fmt"
	"grover/code"
)

type Item struct {
	Functions []Function
}

type Function struct {
	Name string
}

func Parse(file string) []Item {
	items := []Item{}

	lines := code.ReadFileRemoveNewlines(file)
	fmt.Println(lines)
	return items
}
