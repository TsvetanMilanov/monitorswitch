package main

import (
	"fmt"
	"strconv"
)

// OffCommand type
type OffCommand struct {
}

// Execute executes the command with the provided flags
func (command *OffCommand) Execute(commandArguments *CommandArguments) {
	fmt.Println("Switching monitor off...")

	var monitorsService = injector.monitorsService
	var indexParameter = commandArguments.commandParameters[0]
	var monitorIndex, err = strconv.Atoi(indexParameter)

	if err != nil {
		panic(fmt.Sprintf("Index %v is not a valid value for monitor.\n", indexParameter))
	}

	monitorIndex--

	var allMonitors = monitorsService.GetAllMonitors()

	var monitor = monitorsService.GetMonitor(allMonitors, commandArguments.commandFlags.primary)
	if commandArguments.commandFlags.primary {
		monitor = monitorsService.GetMonitor(allMonitors, commandArguments.commandFlags.primary)
	} else {
		monitor = allMonitors[monitorIndex]
	}

	monitorsService.SwitchMonitorOff(monitor)
	fmt.Printf("Monitor %s switched off.\n", monitor.name)
}

// GetAliases returns the aliases of the command
func (command *OffCommand) GetAliases() []string {
	return []string{"off"}
}
