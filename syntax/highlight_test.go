package syntax

import (
	"strings"
	"testing"
)

func TestHighlight(t *testing.T) {
	line := "foo\t\tbar  and     spaces."
	test1 := Highlight(line)
	tokens := strings.Split(test1, " ")
	if len(tokens) != 12 {
		t.Fatal("wrong length")
	}
}
