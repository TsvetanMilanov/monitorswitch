package main

import "fmt"

// ListMonitorsCommand type
type ListMonitorsCommand struct {
}

// Execute executes the command with the provided flags
func (command *ListMonitorsCommand) Execute(commandArguments *CommandArguments) {
	fmt.Println("Monitors:")
	var monitorsService = injector.monitorsService
	var monitors = monitorsService.GetAllMonitors()

	for i, monitor := range monitors {
		fmt.Printf("%d - %s primary: %v\n", i+1, monitor.name, monitor.isPrimary)
	}
}

// GetAliases returns the aliases of the command
func (command *ListMonitorsCommand) GetAliases() []string {
	return []string{"list-monitors", "ls"}
}
