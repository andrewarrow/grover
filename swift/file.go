package swift

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(path string) {
	asBytes, _ := ioutil.ReadFile(path)
	asString := string(asBytes)
	for _, c := range asString {
		char := fmt.Sprintf("%c", c)
		fmt.Printf(char)
	}
}
