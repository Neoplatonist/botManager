package modules

import "strings"

var registered, connected []string

func addRegister(m string) {
	registered = append(registered, m)
}

func rmRegister(m string) {
	for i, module := range registered {
		if module == m {
			registered = append(registered[:i], registered[i+1:]...)
		}
	}
}

func addConnected(m string) {
	connected = append(connected, m)
}

func rmConnected(m string) {
	for i, module := range connected {
		if module == m {
			connected = append(connected[:i], connected[i+1:]...)
		}
	}
}

// RegisteredModules returns all modules currently registered
func RegisteredModules() string {
	return strings.Join(registered, "\n")
}

// ConnectedModules returns all modules currently in use
func ConnectedModules() string {
	return strings.Join(connected, "\n")
}
