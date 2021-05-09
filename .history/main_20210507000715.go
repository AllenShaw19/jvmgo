package main

import (
	"fmt"

	"github.com/jvmgo/cmd"
)

const version = "0.0.1"

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Printf("version %s\n", version)
	} else if command.HelpFlag || command.Class == "" {
		cmd.PrintUsage()
	} else {

	}
}

func sta