package main

import (
	"fmt"
	"github.com/jvmgo/cmd"
)

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println()
	}
}
