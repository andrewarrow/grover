package swift

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	items := Parse("example.swift")
	fmt.Println(items)
}
