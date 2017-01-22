package main

import (
	"fmt"
)

// ExecuteCommand executes command.
func ExecuteCommand(commandArguments *CommandArguments) {
	var commandsRegistry = BuildCommandRegistry()

	for _, command := range *commandsRegistry {
		var aliases = command.GetAliases()
		for _, alias := range aliases {
			if alias == commandArguments.commandName {
				command.Execute(commandArguments)
				return
			}
		}
	}

	panic(fmt.Sprintf("Unknown command %s", commandArguments.commandName))
}
