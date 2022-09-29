package main

import (
	"grover/files"
	"grover/scan"
	"grover/screen"
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
	files.Mkdir("data")
	arg1 := os.Args[1]

	if arg1 == "help" {
	} else if arg1 == "scan" {
		dir := util.GetArg(2)
		filter := util.GetArg(3)
		scan.Scan(dir, filter)
	} else {
		path := screen.NewPath(arg1)
		screen.OnePath(path)
	}
}
