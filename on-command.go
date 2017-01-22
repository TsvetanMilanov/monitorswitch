package main

import (
	"fmt"
)

// OnCommand type
type OnCommand struct {
}

// Execute executes the command with the provided flags
func (command *OnCommand) Execute(commandFlags *CommandFlags) {
	fmt.Println("Switching monitor on...")

	var monitorsService = injector.monitorsService

	var allMonitors = monitorsService.GetAllMonitors()

	var monitor = monitorsService.GetMonitor(allMonitors, commandFlags.primary)
	var referenceMonitor = monitorsService.GetMonitor(allMonitors, !commandFlags.primary)
	var side string

	if commandFlags.left {
		side = "--left-of"
	} else {
		side = "--right-of"
	}

	monitorsService.SwitchMonitorOn(monitor, referenceMonitor, side)
	fmt.Printf("Monitor %s switched on.\n", monitor.name)
}

// GetAliases returns the name of the command
func (command *OnCommand) GetAliases() []string {
	return []string{"on"}
}

func getMonitor(monitors []*Monitor, shouldGetPrimaryMonitor bool) *Monitor {
	// TODO: refactor this to work with monitor name
	for _, monitor := range monitors {
		if monitor.isPrimary {
			return monitor
		}
	}

	return monitors[len(monitors)-1]
}
