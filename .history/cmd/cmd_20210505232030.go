package cmd

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
}

