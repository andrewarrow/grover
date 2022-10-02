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
		for _, withColor := range evalToken(t) {
			if withColor.Color == "" {
				buff = append(buff, withColor.Text)
			} else {
				buff = append(buff, fmt.Sprintf("[%s](fg:%s)", withColor.Text,
					withColor.Color))
			}
		}
	}
	return strings.Join(buff, " ")
}

func evalToken(t string) []*Token {
	color := "white"
	if rand.Intn(2) == 0 {
		color = "green"
	}

	tokens := []*Token{}
	newToken := &Token{}
	newToken.Text = t
	newToken.Color = color

	if t == "" {
		newToken.Color = ""
		tokens = append(tokens, newToken)
		return tokens
	}

	open := strings.Split(t, "(")
	closed := strings.Split(t, ")")

	if len(open) == 2 && len(closed) == 1 {
		newToken := &Token{}
		newToken.Text = open[0]
		newToken.Color = "white"
		tokens = append(tokens, newToken)
		newToken = &Token{}
		newToken.Text = "("
		newToken.Color = "cyan"
		tokens = append(tokens, newToken)
		newToken = &Token{}
		newToken.Text = open[1]
		newToken.Color = "white"
		tokens = append(tokens, newToken)
	} else if len(open) == 1 && len(closed) == 2 {
	} else if len(open) > 1 && len(closed) > 1 {
	} else {
		if t == "func" {
			newToken.Color = "cyan"
		} else if t == "string" {
			newToken.Color = "magenta"
		}
		tokens = append(tokens, newToken)
	}

	// t = func
	// t = foo(bar
	// t = string)
	// t = foo()
	return tokens
}
