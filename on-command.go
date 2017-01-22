package main

import "fmt"
import "strconv"

// OnCommand type
type OnCommand struct {
}

// Execute executes the command with the provided flags
func (command *OnCommand) Execute(commandArguments *CommandArguments) {
	fmt.Println("Switching monitor on...")

	var monitorsService = injector.monitorsService
	var indexParameter = commandArguments.commandParameters[0]
	var monitorIndex, err = strconv.Atoi(indexParameter)

	if err != nil {
		panic(fmt.Sprintf("Index %v is not a valid value for monitor.\n", indexParameter))
	}

	monitorIndex--

	var allMonitors = monitorsService.GetAllMonitors()

	var monitor = monitorsService.GetMonitor(allMonitors, commandArguments.commandFlags.primary)
	var referenceMonitor = monitorsService.GetMonitor(allMonitors, !commandArguments.commandFlags.primary)
	if commandArguments.commandFlags.primary {
		monitor = monitorsService.GetMonitor(allMonitors, commandArguments.commandFlags.primary)
		referenceMonitor = monitorsService.GetMonitor(allMonitors, !commandArguments.commandFlags.primary)
	} else {
		monitor = allMonitors[monitorIndex]
		referenceMonitor = getClosestMonitor(allMonitors, monitorIndex)
	}

	var side string

	if commandArguments.commandFlags.left {
		side = "--left-of"
	} else {
		side = "--right-of"
	}

	monitorsService.SwitchMonitorOn(monitor, referenceMonitor, side)
	fmt.Printf("Monitor %s switched on.\n", monitor.name)
}

// GetAliases returns the aliases of the command
func (command *OnCommand) GetAliases() []string {
	return []string{"on"}
}

func getClosestMonitor(monitors []*Monitor, index int) *Monitor {
	var previousIndex = index - 1
	var nextIndex = index + 1

	if previousIndex >= 0 {
		return monitors[previousIndex]
	}

	return monitors[nextIndex]
}
