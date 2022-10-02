package syntax

import (
	"fmt"
	"math/rand"
	"strings"
)

type Token struct {
	Text  string
	Color string
}

func Highlight(s string) string {
	replaced := strings.Replace(s, "\t", "  ", -1)
	tokens := strings.Split(replaced, " ")
	buff := []string{}
	for _, t := range tokens {
		for _, tokenWithColor := range evalToken(t) {
			buff = append(buff, fmt.Sprintf("[%s](fg:%s)", tokenWithColor.Text,
				tokenWithColor.Color))
		}
	}
	return strings.Join(buff, " ")
}

func evalToken(t string) []*Token {
	tokens := []*Token{}
	color := "white"
	if rand.Intn(2) == 0 {
		color = "green"
	}
	open := strings.Split(t, "(")
	closed := strings.Split(t, ")")
	newToken := &Token{}

	if len(open) > 1 && len(closed) == 1 {
		fmt.Println("1", t)
	} else if len(open) == 1 && len(closed) > 1 {
		fmt.Println("2", t)
	} else if len(open) > 1 && len(closed) > 1 {
		fmt.Println("3", t)
	} else {
		fmt.Println("4", t)
		newToken.Text = t
		newToken.Color = color
		tokens = append(tokens, newToken)
	}

	// t = func
	// t = foo(bar
	// t = string)
	// t = foo()
	return tokens
}
