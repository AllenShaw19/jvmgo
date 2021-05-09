package cmd

import "fmt"

// java [-options] class [args...]
type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
}

func ParseCmd() *Cmd {

}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]")
}