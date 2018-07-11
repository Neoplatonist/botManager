package command

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/now"
	"github.com/mattn/go-shellwords"
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

	ActiveCommands["Neo-Bot"] = []Command{
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
			Name: "!remind",
			Desc: "user in x seconds: !remind <username> \"<message>\" <seconds>",
			Action: func(s *discordgo.Session, m *discordgo.MessageCreate) {
				// cmd := strings.Fields(m.Content)
				cmd, err := shellwords.Parse(m.Content)
				if err != nil {
					fmt.Println(err)
					return
				}

				if m.Author.ID == s.State.User.ID || !strings.EqualFold(cmd[0], "!remind") {
					return
				}

				fmt.Println(cmd)

				if len(cmd) < 4 {
					s.ChannelMessageSend(
						m.ChannelID,
						fmt.Sprintf("%s sorry not enough commands \n!remind <username> <one word message> <seconds>", m.Author.Username),
					)
				}

				// user, err := s.User("@Neo-Bot")
				// if err != nil {
				// 	fmt.Println(err)
				// }

				// guild, err := s.UserGuilds(5, "", "")
				// if err != nil {
				// 	fmt.Println(err)
				// }

				// fmt.Println("\n", guild[0].ID)

				// db.Session.Collection("users").find("*")

				// db.Connect()
				// users := db.Collection("discord-users")
				// users.Find(bson.M{"username": user})

				list, err := s.GuildMembers("406302666050895882", "", 100)
				if err != nil {
					fmt.Println(err)
				}

				var person *discordgo.User
				for _, user := range list {
					if user.User.Username == cmd[1] || user.Nick == cmd[1] {
						fmt.Println(user.User.Username)
						person = user.User
					}
				}

				if person == nil {
					s.ChannelMessageSend(
						m.ChannelID,
						fmt.Sprintf("%v, sorry no users found", m.Author.Username),
					)

					return
				}

				t, _ := strconv.Atoi(cmd[3])
				time.Sleep(time.Duration(t) * time.Second)

				s.ChannelMessageSend(
					m.ChannelID,
					fmt.Sprintf("%v %s", person.Mention(), cmd[2]),
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

// !remind @Neoplatonist message 11:15:20 5
// !help remind
// !remind <username> <message> <hh:mm:ss> (optional)<days>
// !remind <username> <message> hours=1 minutes=1 seconds=1 days=1
