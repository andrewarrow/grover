package code

import (
	"fmt"
	"strings"
)

func ProcessCharacters(characters []string) {
	word := []string{}
	lastWords := []string{}
	spaces := ""
	for _, c := range characters {
		if c == "{" {
			fmt.Printf("%s%s\n", spaces, lastWords[len(lastWords)-1])
			spaces = spaces + "  "
		} else if c == "}" {
			spaces = spaces[0 : len(spaces)-2]
		} else if c == " " {
			completeWord := strings.Join(word, "")
			lastWords = append(lastWords, completeWord)
			word = []string{}
		} else {
			word = append(word, c)
		}
	}

}
