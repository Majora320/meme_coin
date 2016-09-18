package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "database/sql"
	_ "strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"

	_ "github.com/bmizerany/pq"
)

// Variables used for command line parameters
var (
	Token string
	Email string
	PW    string
	BotID string
)

type User struct {
	ID        string `db:"id"`
	Username  string `db:"name"`
	CurMoney  int    `db:"current_money"`
	TotMoney  int    `db:"total_money"`
	WonMoney  int    `db:"won_money"`
	LostMoney int    `db:"lost_money"`
	GiveMoney int    `db:"given_money"`
	RecMoney  int    `db:"received_money"`
}

func db_get() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost user=memebot dbname=money password=password sslmode=disable parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func init() {
	Token, _ = os.LookupEnv("bot_token")
	Email, _ = os.LookupEnv("email")
	PW, _ = os.LookupEnv("pw")

}

func main() {

	// Create a new Discord session using the provided login information.
	dg, err := discordgo.New(Email, PW)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	BotID = u.ID

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running!")
	<-make(chan struct{})
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if strings.Contains(m.Content, "!tip") == true {
		args := strings.Split(m.Content, " ")
		if len(args) > 3 {
			amount := args[1]
			from := m.Author
			for _, to := range m.Mentions {
				_, _ = s.ChannelMessageSend(m.ChannelID, "tip "+amount+" dankmemes to "+to.Username+" from: "+from.Username)
			}
		} else {
			return
		}
	}

	if m.Content == "meme" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "you're a memestar harry")
	}
}
