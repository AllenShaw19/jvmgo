package main

import (
	"fmt"

	"github.com/jvmgo/cmd"
)

const version = "0.0.1"

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Printf("version %s", version)
	}
}
