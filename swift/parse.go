package swift

import (
	"grover/code"
	"strings"
)

type SwiftFile struct {
	Classes []Class
}

type Class struct {
	Name      string
	Functions []Function
}

type Function struct {
	Name string
}

func Parse(file string) SwiftFile {
	item := SwiftFile{}

	chars := code.ReadFileRemoveNewlines(file)
	words := ProcessCharacters(chars)
	for _, c := range findNext("class", words) {
		item.Classes = append(item.Classes, Class{c, []Function{}})
	}
	return item
}

func findNext(match string, words []string) []string {
	items := []string{}
	for i, w := range words {
		if w == "class" {
			items = append(items, words[i+1])
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
