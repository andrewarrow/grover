package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ProduceFakeProject() {
	RmRfBang()
	WriteFile("BigCompany.xcodeproj", "")
}

func WriteFile(path, content string) {
	fname := fmt.Sprintf("%s/%s", "data", path)
	os.Remove(fname)
	ioutil.WriteFile(fname, []byte(content), 0644)
}
