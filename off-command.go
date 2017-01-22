package main

import (
	"fmt"
)

// OffCommand type
type OffCommand struct {
}

// Execute executes the command with the provided flags
func (command *OffCommand) Execute(commandFlags *CommandFlags) {
	fmt.Println("Switching monitor off...")

	var monitorsService = injector.monitorsService

	var allMonitors = monitorsService.GetAllMonitors()

	var monitor = monitorsService.GetMonitor(allMonitors, commandFlags.primary)

	monitorsService.SwitchMonitorOff(monitor)
	fmt.Printf("Monitor %s switched off.\n", monitor.name)
}

// GetAliases returns the name of the command
func (command *OffCommand) GetAliases() []string {
	return []string{"off"}
}
