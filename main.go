package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"unicode"

	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Core struct {
	bot  *discordgo.Session
	spam *string
}

var core Core

func main() {
	EnvInit()
	GNULinuxBible()
	DiscordInit()

	fmt.Println("Spam cached:")
	fmt.Println(*core.spam)
	fmt.Println("Opening bot connection...")

	core.bot.AddHandler(ListenLinuxMsg)

	err := core.bot.Open()
	defer core.bot.Close()

	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Closing bot connection...")
}

func EnvInit() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GNULinuxBible() {
	response, err := http.Get("https://stallman-copypasta.github.io")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer response.Body.Close()
	spmm := ""
	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		spmm = formatBible(string(body))
	} else {
		spmm = "XD"
	}

	core.spam = &spmm
}

func formatBible(str string) string {
	// From-to main tag
	start := strings.Index(str, "<main>")
	end := strings.Index(str, "</main>")
	str = str[start:end]

	ignoreTag := false
	content := ""
	for _, char := range str {
		// ignore all html tags
		if char == '<' {
			ignoreTag = true
			continue
		} else if char == '>' {
			ignoreTag = false
			continue
		}

		// remove any non-graphic rune except newlines
		if ignoreTag || (char != '\n' && !unicode.IsGraphic(char)) {
			continue
		}

		// concat it
		content += string(char)
	}
	return content
}

func DiscordInit() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))

	if err != nil {
		log.Fatal(err)
		return
	}
	core.bot = discord
}

func ListenLinuxMsg(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	if strings.Contains(strings.ToLower(message.Content), "linux") {
		_, e := discord.ChannelMessageSend(message.ChannelID, *core.spam)
		if e != nil {
			log.Fatal(e)
		}
	}
}
