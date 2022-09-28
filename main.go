package main

import (
	"grover/scan"
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
		scan.Scan(os.Args[2])
	}
}
