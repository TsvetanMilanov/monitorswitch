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

	if commandFlags.right {
		side = "--right-of"
	} else {
		side = "--left-of"
	}

	monitorsService.SwitchMonitorOn(monitor, referenceMonitor, side)
	fmt.Printf("Monitor %s switched on.\n", monitor.name)
}

// GetName returns the name of the command
func (command *OnCommand) GetName() string {
	return "on"
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
