package gengine

import "strconv"

type Command struct {
	Name        string
	Description string
	Args        Args
	Run         func([]string)
}

type Args struct {
	Count int
	Get   []string
}

func RegisterCommand(name string, description string, args []string, runnable func([]string)) {
	arguments := Args{
		Count: len(args),
		Get:   args,
	}
	commands = append(commands, Command{
		Name:        name,
		Description: description,
		Args:        arguments,
		Run:         runnable,
	})

	Print(Prefix(0) + "Successfully registered command \"" + name + "\" with " + strconv.Itoa(len(args)) + " arguments.")
}
