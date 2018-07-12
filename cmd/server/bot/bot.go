package bot

import (
	"fmt"

	"github.com/neoplatonist/botManager/cmd/server/modules"
)

var ModuleList = map[string]module{
	"Neo-Bot": modules.Discord(),
}

type module interface {
	Register() error
	Connect() error
	Disconnect() error
	Name() string
	Command([]string) (string, error)
}

func initModules() {
	for _, module := range ModuleList {
		if err := module.Register(); err != nil {
			fmt.Printf("could not register module: %s", err)
		}
	}

	fmt.Printf("Modules Registered: \n%s\n", modules.ActiveModules())
}

func connectModules() error {
	for _, module := range ModuleList {
		if err := module.Connect(); err != nil {
			return fmt.Errorf("could not connect to module: %s", err)
		}
	}

	return nil
}

// Connect individual modules
func Connect(name string) (string, error) {
	if err := ModuleList[name].Connect(); err != nil {
		return "", err
	}

	return ModuleList[name].Name() + " has connected", nil
}

// Disconnect individual modules
func Disconnect(name string) (string, error) {
	err := ModuleList[name].Disconnect()
	if err != nil {
		return "", err
	}

	return ModuleList[name].Name() + " has disconnected", nil
}

// Start initializes bots
func Start() {
	initModules()
	connectModules()
}
