package gengine

import (
	"bufio"
	"os"
	"strings"
)

var (
	scanner = bufio.NewScanner(os.Stdin)
)

func ValueOf(thing string, err error) string {
	if err != nil {
		Print(Prefix(0) + "An error occoured: " + err.Error())
	}
	return thing
}

func RemoveElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func Input(msg string) string {
	PrintR(msg)
	scanner.Scan()
	return scanner.Text()
}

func RInput(msg string) string {
	PrintR(msg)
	scanner.Scan()
	return scanner.Text()
}

func CommandInput() (string, []string) {
	commandLine := commandline

	userInput := Input(commandLine)
	splitted := strings.Split(userInput, " ")
	command := splitted[0]
	args := RemoveElement(splitted, 0)
	return command, args
}
