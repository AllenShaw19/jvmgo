package main

import (
	"fmt"
	"strings"

	"github.com/jvmgo/classpath"
	"github.com/jvmgo/cmd"
	"github.com/jvmgo/logger"
)

const version = "0.0.1"

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Printf("version %s\n", version)
	} else if command.HelpFlag || command.Class == "" {
		cmd.PrintUsage()
	} else {
		startJvm(command)
	}
}

func startJvm(cmd *cmd.Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	logger.Info("classpath:%v, class:%v, args:%v\n", cp, cmd.Class, cmd.Args)

	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		logger.
	}

}
