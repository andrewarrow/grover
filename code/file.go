package code

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) []string {
	asBytes, _ := ioutil.ReadFile(path)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")

	buff := []string{}
	for _, line := range asLines {
		trimmed := strings.TrimSpace(line)
		buff = append(buff, trimmed)
	}
	return buff
}

func ReadFileRemoveNewlines(path string) []string {
	asBytes, _ := ioutil.ReadFile(path)
	asString := string(asBytes)
	asLines := strings.Split(asString, "\n")

	buff := []string{}
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
			buff = append(buff, char)
		}
		buff = append(buff, " ")
	}
	return buff
}
