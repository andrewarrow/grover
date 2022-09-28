package syntax

import (
	"fmt"
	"math/rand"
	"strings"
)

func Highlight(s string) string {
	tokens := strings.Split(s, " ")
	buff := []string{}
	for _, t := range tokens {
		color := "white"
		if rand.Intn(2) == 0 {
			color = "green"
		}
		buff = append(buff, fmt.Sprintf("[%s](fg:%s)", t, color))
	}
	return strings.Join(buff, " ")
}
