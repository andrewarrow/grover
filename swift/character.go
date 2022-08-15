package swift

import (
	"fmt"
	"strings"
)

func ProcessCharacters(characters []string) {
	word := []string{}
	for _, c := range characters {
		if c == " " {
			fmt.Println(strings.Join(word, ""))
			word = []string{}
		} else {
			word = append(word, c)
		}
	}

}
