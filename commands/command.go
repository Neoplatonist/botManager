package command

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/leekchan/timeutil"
)

var ActiveCommands []Command

// Command defines the database structure of all commands
type Command struct {
	Name   string
	Desc   string
	Action func(s *discordgo.Session, m *discordgo.MessageCreate)
}

// GetList searches the database for all commands for a given plugin.
// The plugin is received by a string
func GetList(s string) ([]Command, error) {
	// connect to a database
	// search for plugin by string

	ActiveCommands = []Command{
		Command{
			Name: "!time",
			Desc: "Time until Neoplatonist's next stream.",
			Action: func(s *discordgo.Session, m *discordgo.MessageCreate) {
				if m.Author.ID == s.State.User.ID || !strings.EqualFold(m.Content, "!time") {
					return
				}

				type tyme struct {
					hour   int
					minute int
					second int
				}

				type date struct {
					tyme
					stamp time.Time
				}

				type env struct {
					day      int
					current  date
					delta    tyme
					schedule map[int]tyme
				}

				var err error
				e := &env{}
				e.schedule = map[int]tyme{
					2: tyme{10, 0, 0},
					3: tyme{10, 0, 0},
					4: tyme{10, 0, 0},
					5: tyme{10, 0, 0},
				}

				flatten := func() {
					if e.delta.second < 0 {
						e.delta.minute--
						e.delta.second += 60
					}

					if e.delta.minute < 0 {
						e.delta.hour--
						e.delta.minute += 60
					}
				}

				var calcHours func(int) int
				calcHours = func(diffHour int) int {
					if _, ok := e.schedule[e.day+1]; !ok {
						if e.day > 5 {
							// Restart the week
							e.day = 0
							diffHour = calcHours(diffHour + 24)
						} else {
							// Scoot up to the next scheduled day
							e.day++
							diffHour = calcHours(diffHour + 24)
						}
					} else {
						// Calculate if only next day
						diffHour = diffHour + 24 - 12
					}

					return diffHour
				}

				calcDifference := func() string {
					var err error

					e.day, err = strconv.Atoi(timeutil.Strftime(&e.current.stamp, "%w"))
					if err != nil {
						fmt.Printf("could not parse string to int: %s", err)
					}

					e.delta.hour = e.schedule[e.day].hour - e.current.hour
					e.delta.minute = e.schedule[e.day].minute - e.current.minute
					e.delta.second = e.schedule[e.day].second - e.current.second

					flatten()

					if e.delta.hour < 0 {
						e.delta.hour = calcHours(e.current.hour)

						e.delta.minute = e.schedule[e.day].minute - e.current.minute
						e.delta.second = e.schedule[e.day].second - e.current.second

						flatten()
					}

					return fmt.Sprintf("%02d:%02d:%02d", e.delta.hour, e.delta.minute, e.delta.second)
				}

				e.current.stamp = time.Now()
				e.current.hour, err = strconv.Atoi(timeutil.Strftime(&e.current.stamp, "%H"))
				if err != nil {
					fmt.Printf("could not parse hour from timestamp: %s", err)
				}

				e.current.minute, err = strconv.Atoi(timeutil.Strftime(&e.current.stamp, "%M"))
				if err != nil {
					fmt.Printf("could not parse minute from timestamp: %s", err)
				}

				e.current.second, err = strconv.Atoi(timeutil.Strftime(&e.current.stamp, "%S"))
				if err != nil {
					fmt.Printf("could not parse second from timestamp: %s", err)
				}

				var greeting string
				if 5 <= e.current.hour && e.current.hour < 11 {
					greeting = "Ohayou Gozaimasu"
				} else if 11 <= e.current.hour && e.current.hour < 18 {
					greeting = "Konnichiwa"
				} else {
					greeting = "Konbanwa"
				}

				result := calcDifference()

				s.ChannelMessageSend(
					m.ChannelID,
					fmt.Sprintf(
						"%s %v, from Japan! \nTime until next stream: %s",
						greeting,
						"Neoplatonist",
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
					fmt.Sprintf("List of commands: \n%s", CommandList()),
				)
			},
		},
	}

	for _, command := range ActiveCommands {
		cmd := command.Name + " - " + command.Desc
		addCommand(cmd)
	}

	return ActiveCommands, nil
}
