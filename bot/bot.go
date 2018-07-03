package bot

import (
	"fmt"

	"github.com/neoplatonist/discord-bot/bot-v2/modules"
)

var moduleList = []module{
	modules.Discord(),
}

type module interface {
	Register() error
	Connect() error
}

func initModules() {
	for _, module := range moduleList {
		if err := module.Register(); err != nil {
			fmt.Printf("could not register plugin: %s", err)
		}
	}

	fmt.Printf("Modules Registered: \n%s\n", modules.ActiveModules())
	fmt.Println("------------------------------")
}

func connectModules() {
	for _, module := range moduleList {
		if err := module.Connect(); err != nil {
			fmt.Printf("could not connect to plugin: %s", err)
		}
	}
}

// Start initializes bots
func Start() {
	initModules()
	connectModules()
}
