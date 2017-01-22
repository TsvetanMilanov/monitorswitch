package main

// Command interface
type Command interface {
	GetName() string
	Execute(commandFlags *CommandFlags)
}
