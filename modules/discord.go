package modules

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/neoplatonist/botManager/commands"
)

// State instantiates the session state
var State DiscordState

type DiscordState struct {
	Session *discordgo.Session
}

func Discord() DiscordState {
	return DiscordState{}
}

func (d DiscordState) Register() error {
	session, err := discordgo.New("Bot " + os.Getenv("DISCORD_APP_USER"))
	if err != nil {
		return fmt.Errorf("could not create a Discord session: %s", err)
	}

	commandsList, err := command.GetList("discord")
	if err != nil {
		return fmt.Errorf("no discord commands found: %s", err)
	}

	for _, command := range commandsList {
		session.AddHandler(command.Action) // Tenative
	}

	addModule("discord-bot")
	State = DiscordState{session}

	return nil
}

func (d DiscordState) Connect() error {
	if err := State.Session.Open(); err != nil {
		return fmt.Errorf("could not open discord connection: %s", err)
	}

	fmt.Println("Neo-Bot is now running")
	return nil
}

func (d DiscordState) Disconnect() error {
	if err := State.Session.Close(); err != nil {
		return fmt.Errorf("could not close discord connection: %s", err)
	}

	fmt.Println("Neo-Bot is now offline")
	return nil
}
