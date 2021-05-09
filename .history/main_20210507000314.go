package main

import "github.com/jvmgo/cmd"

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Pr
	}
}
