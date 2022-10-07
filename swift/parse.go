package swift

import (
	"grover/code"
	"strings"
)

func Parse(file string) Block {
	item := NewBlock("")

	chars := code.ReadFileRemoveNewlines(file)
	findBlocks(chars, &item)
	item.printBlocks("")
	/*
		words := ProcessCharacters(chars)
		for _, c := range findNext("class", words) {
			item.Classes = append(item.Classes, NewClass(c))
		}*/
	return item
}

func findBlocks(characters []string, block *Block) {
	text := []string{}
	stack := []*Block{block}
	for _, c := range characters {
		if c == "{" {
			b := NewBlock(strings.Join(text, ""))
			endOf(stack).Blocks = append(endOf(stack).Blocks, &b)
			stack = append(stack, &b)
			text = []string{}
		} else if c == "}" {
			stack = stack[0 : len(stack)-1]
		} else {
			text = append(text, c)
		}
	}
}

func endOf(stack []*Block) *Block {
	return stack[len(stack)-1]
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
