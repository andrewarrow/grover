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

	if len(open) > 0 && len(closed) == 0 {
	} else if len(closed) > 0 && len(open) == 0 {
	} else if len(closed) > 0 && len(open) > 0 {
	} else {
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
