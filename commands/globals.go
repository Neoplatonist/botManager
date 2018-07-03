package command

import "strings"

var commandList []string

func addCommand(c string) {
	commandList = append(commandList, c)
}

// CommandList returns all commands currently usable
func CommandList() string {
	return strings.Join(commandList, "\n")
}
