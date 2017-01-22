package main

// Command interface
type Command interface {
	GetAliases() []string
	Execute(commandArguments *CommandArguments)
}
