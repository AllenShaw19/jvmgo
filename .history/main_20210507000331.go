package main

import (
	"fmt"
	"github.com/jvmgo/cmd"
)

const 

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println()
	}
}
