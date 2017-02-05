package utils

import (
	"strings"

	"github.com/TsvetanMilanov/monitorswitch/commands"
)

// ParseCommandArguments returns the command flags
func ParseCommandArguments(arguments []string) *commands.CommandArguments {
	var commandFlags = new(commands.CommandFlags)
	var commandArguments = new(commands.CommandArguments)

	// The first command argument is the binary name
	for i := 1; i < len(arguments); i++ {
		var commandArgument = arguments[i]

		if commandArgument[0] == '-' {
			// Command flag
			var flagName = strings.TrimLeft(commandArgument, "-")
			switch flagName {
			case "primary":
				commandFlags.Primary = true
				break
			case "left":
				commandFlags.Left = true
				break
			case "right":
				commandFlags.Right = true
				break
			}
		} else if i <= 1 {
			// Command name
			commandArguments.CommandName = commandArgument
		} else {
			// Command parameter
			commandArguments.CommandParameters = append(commandArguments.CommandParameters, commandArgument)
		}
	}

	commandArguments.CommandFlags = commandFlags

	return commandArguments
}
