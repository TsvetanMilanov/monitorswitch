package commands

// CommandArguments type
type CommandArguments struct {
	CommandName       string
	CommandParameters []string
	CommandFlags      *CommandFlags
}
