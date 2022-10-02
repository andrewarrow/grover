package syntax

import (
	"fmt"
	"strings"
	"testing"
)

func TestHighlight(t *testing.T) {
	line := "foo\t\tbar  (and  more)  foo()   spaces."
	test := Highlight(line)
	fmt.Println(len(test), test)
	tokens := strings.Split(test, " ")
	if len(tokens) != 12 {
		t.Fatal("wrong length")
	}
}
