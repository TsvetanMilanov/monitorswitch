package commands

import (
	"fmt"
	"strconv"

	"github.com/TsvetanMilanov/monitorswitch/globals"
	"github.com/TsvetanMilanov/monitorswitch/models"
)

// OnCommand type
type OnCommand struct {
}

// Execute executes the command with the provided flags
func (command *OnCommand) Execute(commandArguments *CommandArguments) {
	fmt.Println("Switching monitor on...")

	var monitorsService = globals.GetInjector().MonitorsService
	var indexParameter = commandArguments.CommandParameters[0]
	var monitorIndex, err = strconv.Atoi(indexParameter)

	if err != nil {
		panic(fmt.Sprintf("Index %v is not a valid value for monitor.\n", indexParameter))
	}

	monitorIndex--

	var allMonitors = monitorsService.GetAllMonitors()

	var monitor = monitorsService.GetMonitor(allMonitors, commandArguments.CommandFlags.Primary)
	var referenceMonitor = monitorsService.GetMonitor(allMonitors, !commandArguments.CommandFlags.Primary)
	if commandArguments.CommandFlags.Primary {
		monitor = monitorsService.GetMonitor(allMonitors, commandArguments.CommandFlags.Primary)
		referenceMonitor = monitorsService.GetMonitor(allMonitors, !commandArguments.CommandFlags.Primary)
	} else {
		monitor = allMonitors[monitorIndex]
		referenceMonitor = getClosestMonitor(allMonitors, monitorIndex)
	}

	var side string

	if commandArguments.CommandFlags.Left {
		side = "--left-of"
	} else {
		side = "--right-of"
	}

	monitorsService.SwitchMonitorOn(monitor, referenceMonitor, side)
	fmt.Printf("Monitor %s switched on.\n", monitor.Name)
}

// GetAliases returns the aliases of the command
func (command *OnCommand) GetAliases() []string {
	return []string{"on"}
}

func getClosestMonitor(monitors []*models.Monitor, index int) *models.Monitor {
	var previousIndex = index - 1
	var nextIndex = index + 1

	if previousIndex >= 0 {
		return monitors[previousIndex]
	}

	return monitors[nextIndex]
}
