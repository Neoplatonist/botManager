package modules

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/neoplatonist/botManager/services/commands"
)

// State instantiates the session state
var State DiscordState

// DiscordState contains the discord session state
type DiscordState struct {
	Session *discordgo.Session
}

// Discord initializes the struct
func Discord() DiscordState {
	return DiscordState{}
}

// Register creates a session and connects the command list
func (d DiscordState) Register() error {
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_APP_USER"))
	if err != nil {
		return fmt.Errorf("could not create a Discord session: %s", err)
	}

	moduleName := d.Name()

	commandsList, err := command.GetList(moduleName)
	if err != nil {
		return fmt.Errorf("no discord commands found: %s", err)
	}

	for _, command := range commandsList {
		session.AddHandler(command.Action) // Tenative
	}

	addRegister(moduleName)
	State = DiscordState{session}

	return nil
}

// Connect opens the session
func (d DiscordState) Connect() error {
	moduleName := d.Name()

	if !moduleRegistered(moduleName) {
		return fmt.Errorf("module is not registered: %s", moduleName)
	}

	if err := State.Session.Open(); err != nil {
		return fmt.Errorf("could not open discord connection: %s", err)
	}

	addConnected(d.Name())
	return nil
}

// Disconnect closes the session
func (d DiscordState) Disconnect() error {
	if err := State.Session.Close(); err != nil {
		return fmt.Errorf("could not close discord connection: %s", err)
	}

	rmConnected(d.Name())
	return nil
}

// Name returns the module name
func (d DiscordState) Name() string {
	return "Neo-Bot"
}

// Command allows public commands to be accessed
func (d DiscordState) Command(input []string) (string, error) {
	if len(input) > 1 {
		switch input[1] {
		case "-help":
			list := command.List(input[0])
			return list, nil
		}
	}

	return "", nil
}
