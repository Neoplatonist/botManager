package command

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/now"
)

// ActiveCommands lists all the commands with their actions
var ActiveCommands = make(map[string][]Command, 0)

// Command defines the database structure of all commands
type Command struct {
	Name   string
	Desc   string
	Action func(s *discordgo.Session, m *discordgo.MessageCreate)
}

// GetList searches the database for all commands for a given plugin.
// The plugin is received by a string
func GetList(name string) ([]Command, error) {
	// connect to a database
	// search for plugin by string
	// add commands to list

	ActiveCommands["discord"] = []Command{
		Command{
			Name: "!time",
			Desc: "Time until Neoplatonist's next stream.",
			Action: func(s *discordgo.Session, m *discordgo.MessageCreate) {
				if m.Author.ID == s.State.User.ID || !strings.EqualFold(m.Content, "!time") {
					return
				}

				type env struct {
					day       int
					dayChange int
					current   time.Time
					parsed    time.Time
					duration  time.Duration
					schedule  map[int]string
				}

				e := env{}
				e.schedule = map[int]string{
					2: "10:00:00",
					3: "10:00:00",
					4: "10:00:00",
					5: "10:00:00",
				}

				fmtDuration := func(d time.Duration) string {
					d = d.Round(time.Minute)
					h := d / time.Hour
					d -= h * time.Hour
					m := d / time.Minute
					return fmt.Sprintf("%02d hours and %02d minutes", h, m)
				}

				var calcDiff func() time.Duration
				calcDiff = func() time.Duration {
					if _, ok := e.schedule[e.day]; !ok {
						e.day++
						e.dayChange++
						e.duration = calcDiff()
					}

					var err error
					e.parsed, err = now.Parse(e.schedule[e.day])
					if err != nil {
						fmt.Println(err)
					}

					e.parsed = e.parsed.AddDate(0, 0, e.dayChange)
					e.duration = time.Until(e.parsed)

					if e.duration.Hours() < 0 {
						e.day++
						e.dayChange++
						e.duration = calcDiff()
					}

					return e.duration
				}

				e.current = time.Now()
				e.day = int(e.current.Weekday())
				e.duration = calcDiff()
				result := fmtDuration(e.duration)

				var greeting string
				if 5 <= e.current.Hour() && e.current.Hour() < 11 {
					greeting = "Ohayou Gozaimasu"
				} else if 11 <= e.current.Hour() && e.current.Hour() < 18 {
					greeting = "Konnichiwa"
				} else {
					greeting = "Konbanwa"
				}

				s.ChannelMessageSend(
					m.ChannelID,
					fmt.Sprintf(
						"%s %v, from Japan! \nTime until next stream: %s",
						greeting,
						m.Author.Username,
						result,
					),
				)
			},
		},
		Command{
			Name: "!github",
			Desc: "Returns Neoplatonist's GitHub link.",
			Action: func(s *discordgo.Session, m *discordgo.MessageCreate) {
				if m.Author.ID == s.State.User.ID || !strings.EqualFold(m.Content, "!github") {
					return
				}

				s.ChannelMessageSend(
					m.ChannelID,
					fmt.Sprintf("https://github.com/neoplatonist"),
				)
			},
		},
		Command{
			Name: "!help",
			Desc: "List of all discord commands.",
			Action: func(s *discordgo.Session, m *discordgo.MessageCreate) {
				if m.Author.ID == s.State.User.ID || !strings.EqualFold(m.Content, "!help") {
					return
				}

				s.ChannelMessageSend(
					m.ChannelID,
					fmt.Sprintf("List of commands: \n%s", List("discord")),
				)
			},
		},
	}

	for _, command := range ActiveCommands[name] {
		cmd := command.Name + " - " + command.Desc
		addCommand(name, cmd)
	}

	return ActiveCommands[name], nil
}
