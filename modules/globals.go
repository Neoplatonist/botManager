package modules

import "strings"

var moduleList []string

func addModule(m string) {
	moduleList = append(moduleList, m)
}

// ActiveModules returns all modules currently in use
func ActiveModules() string {
	return strings.Join(moduleList, "\n")
}
