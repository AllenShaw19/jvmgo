package cmd

import (
	"fmt"
	"os"
)

// java [-options] class [args...]
type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
}

func ParseCmd() *Cmd {
	cmd 
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}