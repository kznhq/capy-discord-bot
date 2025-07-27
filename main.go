package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"math/rand"
	"time"

	"github.com/kznhq/capyDiscordBot/handlers"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("Loading .env")

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

	// seeding for the random numbers used to make a random color for new roles
	rand.Seed(time.Now().UnixNano())

	// add your handlers (newMessage takes care of commands and calls the correct response function
	client.AddHandler(handlers.NewMessageHandler)
	client.AddHandler(handlers.React4roleReactionAddHandler)
	client.AddHandler(handlers.React4roleReactionRemoveHandler)

	// open the client
	err = client.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err) 
	} else {
		fmt.Println("Connection opened")
	}
	defer client.Close()

	// allow ctrl c to do ctrl c things
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
