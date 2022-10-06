package swift

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	item := Parse("example.swift")
	fmt.Println(item)
}
