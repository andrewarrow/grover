package syntax

import (
	"fmt"
	"strings"
	"testing"
)

func TestHighlight(t *testing.T) {
	//line := "\t\tfunc Highlight(s string)  string {"
	//line := "if !((something && (this || that)) && more)"
	line := "} else if char == \"[\" {"

	tokens := strings.Split(line, " ")
	for _, t := range tokens {
		if t == "" || strings.Contains(t, "[") || strings.Contains(t, "]") {
			fmt.Println(t)
		} else {
			for _, r := range evalToken(t) {
				//fmt.Println(r.Text, r.Color)
				f := fmt.Sprintf("[%s](fg:%s)", r.Text,
					r.Color)
				fmt.Println(f)
			}
		}
	}
	/*
		test := Highlight(line)
		for _, item := range test {
			fmt.Println("1", item)
		}
		tokens := strings.Split(test, " ")
		if len(tokens) != 12 {
			t.Fatal("wrong length")
		}*/
}
