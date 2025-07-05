package main

import (
	// "bufio"
	// "log"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if m.Content == "!pet" {
		s.ChannelMessageSend(m.ChannelID, "*hapy capy noises*")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env: ", err) 
	}

	discordToken := os.Getenv("DISCORD_TOKEN")

	client, err := discordgo.New("Bot " + discordToken)

	client.AddHandler(messageCreate)

	err = client.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err) 
	}
	defer client.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

}
