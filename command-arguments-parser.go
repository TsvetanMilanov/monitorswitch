package main

import "strings"

// ParseCommandArguments returns the command flags
func ParseCommandArguments(arguments []string) *CommandArguments {
	var commandFlags = new(CommandFlags)
	var commandArguments = new(CommandArguments)

	// The first command argument is the binary name
	for i := 1; i < len(arguments); i++ {
		var commandArgument = arguments[i]

		if commandArgument[0] == '-' {
			// Command flag
			var flagName = strings.TrimLeft(commandArgument, "-")
			switch flagName {
			case "primary":
				commandFlags.primary = true
				break
			case "left":
				commandFlags.left = true
				break
			case "right":
				commandFlags.right = true
				break
			}
		} else if i <= 1 {
			// Command name
			commandArguments.commandName = commandArgument
		} else {
			// Command parameter
			commandArguments.commandParameters = append(commandArguments.commandParameters, commandArgument)
		}
	}

	commandArguments.commandFlags = commandFlags

	return commandArguments
}
