package syntax

import (
	"fmt"
	"testing"
)

func TestHighlight(t *testing.T) {
	test1 := Highlight("foo bar")
	fmt.Println(test1)
}
