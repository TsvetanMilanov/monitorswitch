package main

import (
	"fmt"
)

// ExecuteCommand executes command.
func ExecuteCommand(commandArguments *CommandArguments) {
	var commandsRegistry = BuildCommandRegistry()

	for i := 0; i < len(commandsRegistry); i++ {
		var currentCommand = commandsRegistry[i]

		if currentCommand.GetName() == commandArguments.commandName {
			currentCommand.Execute(commandArguments.commandFlags)
			return
		}
	}

	panic(fmt.Sprintf("Unknown command %s", commandArguments.commandName))
}
