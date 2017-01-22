package main

import "os"

var injector *Injector

func main() {
	injector = bootstrap()
	var commandArguments = ParseCommandArguments(os.Args)

	if len(commandArguments.commandName) == 0 {
		panic("Command name required")
	}

	ExecuteCommand(commandArguments)
}
