package main

import (
	"os"

	"github.com/TsvetanMilanov/monitorswitch/commands"
	"github.com/TsvetanMilanov/monitorswitch/globals"
	"github.com/TsvetanMilanov/monitorswitch/utils"
	"github.com/TsvetanMilanov/monitorswitch/bootstrap"
)

func main() {
	globals.SetupInjector(bootstrap.RunBootstrap)
	var commandArguments = utils.ParseCommandArguments(os.Args)

	if len(commandArguments.CommandName) == 0 {
		panic("Command name required")
	}

	commands.ExecuteCommand(commandArguments)
}
