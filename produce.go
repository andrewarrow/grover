package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ProduceFakeProject() {
	RmRfBang()
	WriteFile("BigCompany.xcodeproj", "")
	WriteFile("BigCompany/AppDelegate.swift", "")
	WriteFile("Things/Foo/Foo/BigCompanyFoo.swift", "")
}

func WriteFile(path, content string) {
	fmt.Println("data/" + ExtractDir(path))
	os.MkdirAll("data/"+ExtractDir(path), 0755)
	fname := fmt.Sprintf("%s/%s", "data", path)
	os.Remove(fname)
	ioutil.WriteFile(fname, []byte(content), 0644)
}

func ExtractDir(path string) string {
	tokens := strings.Split(path, "/")
	buff := []string{}
	for i := 0; i < len(tokens)-1; i++ {
		buff = append(buff, tokens[i])
	}
	return strings.Join(buff, "/")
}
