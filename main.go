package main

import (
	"grover/scan"
	"grover/util"
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
		dir := util.GetArg(2)
		filter := util.GetArg(3)
		scan.Scan(dir, filter)
	}
}
