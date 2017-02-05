package commands

import (
	"fmt"

	"github.com/TsvetanMilanov/monitorswitch/globals"
)

// ListMonitorsCommand type
type ListMonitorsCommand struct {
}

// Execute executes the command with the provided flags
func (command *ListMonitorsCommand) Execute(commandArguments *CommandArguments) {
	fmt.Println("Monitors:")
	var monitorsService = globals.GetInjector().MonitorsService
	var monitors = monitorsService.GetAllMonitors()

	for i, monitor := range monitors {
		fmt.Printf("%d - %s primary: %v\n", i+1, monitor.Name, monitor.IsPrimary)
	}
}

// GetAliases returns the aliases of the command
func (command *ListMonitorsCommand) GetAliases() []string {
	return []string{"list-monitors", "ls"}
}
