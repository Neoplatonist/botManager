package command

import "strings"

var commandList = make(map[string][]string, 0)

func addCommand(name string, c string) {
	commandList[name] = append(commandList[name], c)
}

// List returns all commands currently usable
func List(name string) string {
	return strings.Join(commandList[name], "\n")
}
