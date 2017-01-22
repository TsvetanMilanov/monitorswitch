package main

// BuildCommandRegistry builds the registry with all commands
func BuildCommandRegistry() []Command {
	var result = []Command{new(ListMonitorsCommand),
		new(OnCommand),
		new(OffCommand)}

	return result
}
