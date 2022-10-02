package syntax

import (
	"fmt"
	"strings"
	"testing"
)

func TestHighlight(t *testing.T) {
	line := "\t\tfunc Highlight(s string)  string {"
	test := Highlight(line)
	fmt.Println(len(test), "|"+test+"|")
	tokens := strings.Split(test, " ")
	if len(tokens) != 12 {
		t.Fatal("wrong length")
	}
}
