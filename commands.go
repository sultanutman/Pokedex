package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var pokedexCliCommandsMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Use this command without arguments to exit the program",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Use this command to help you understand how to use a command",
		callback:    commandHelp,
	},
}
