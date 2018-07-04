package bot

import (
	"fmt"

	"github.com/neoplatonist/botManager/modules"
)

var moduleList = map[string]module{
	"discord": modules.Discord(),
}

type module interface {
	Register() error
	Connect() error
	Disconnect() error
}

func initModules() {
	for _, module := range moduleList {
		if err := module.Register(); err != nil {
			fmt.Printf("could not register module: %s", err)
		}
	}

	fmt.Printf("Modules Registered: \n%s\n", modules.ActiveModules())
	fmt.Println("------------------------------")
}

func connectModules() {
	for _, module := range moduleList {
		if err := module.Connect(); err != nil {
			fmt.Printf("could not connect to module: %s", err)
		}
	}
}

// Connect individual modules
func Connect(name string) {
	if err := moduleList[name].Connect(); err != nil {
		fmt.Printf("could not connect to module: %s", err)
	}
}

// Disconnect individual modules
func Disconnect(name string) {
	if err := moduleList[name].Disconnect(); err != nil {
		fmt.Printf("could not disconnect module: %s", err)
	}
}

// Start initializes bots
func Start() {
	initModules()
	connectModules()
}
