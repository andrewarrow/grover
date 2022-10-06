package swift

import (
	"fmt"
	"grover/code"
	"strings"
)

type Item struct {
	Functions []Function
}

type Function struct {
	Name string
}

func Parse(file string) []Item {
	items := []Item{}

	chars := code.ReadFileRemoveNewlines(file)
	words := ProcessCharacters(chars)
	for i, w := range words {
		if w == "func" {
			fmt.Println(words[i+1])
		}
	}
	return items
}

func ProcessCharacters(characters []string) []string {
	word := []string{}
	lastWords := []string{}
	for _, c := range characters {
		if c == "{" {
		} else if c == "}" {
		} else if c == " " {
			completeWord := strings.Join(word, "")
			lastWords = append(lastWords, completeWord)
			word = []string{}
		} else {
			word = append(word, c)
		}
	}
	return lastWords
}
