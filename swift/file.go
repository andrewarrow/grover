package swift

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) {
	asBytes, _ := ioutil.ReadFile(path)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")

	for _, line := range asLines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") {
			continue
		}
		if trimmed == "" {
			continue
		}

		for _, c := range trimmed {
			char := fmt.Sprintf("%c", c)
			fmt.Printf(char)
		}
		fmt.Printf(" ")
	}
}
