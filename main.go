package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	command := os.Args[1]

	if command == "help" {
	} else if command == "scan" {
		Scan(os.Args[2])
	} else if command == "file" {
		OneFile("/Users/andrewarrow/os/vapor/Sources/Vapor/Application.swift")
	}
}
