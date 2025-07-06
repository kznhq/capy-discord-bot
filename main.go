package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

func petCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "!pet" {
		s.ChannelMessageSend(m.ChannelID, "*hapy capy noises*")
	}
}

func main() {
	// need the token to get the client for the bot
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env: ", err) 
	} else {
		fmt.Println(".env loaded")
	}
	discordToken := os.Getenv("DISCORD_TOKEN")

	// make a client
	client, err := discordgo.New("Bot " + discordToken)

	// add your handlers (keeping one per command for cleanliness)
	client.AddHandler(petCommand)

	// open the client
	err = client.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err) 
	} else {
		fmt.Println("Connection opened")
	}
	defer client.Close()

	// ctrl c
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
