package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"

	"github.com/joho/godotenv"
	"github.com/bwmarrin/discordgo"
)

// reads all messages and dispatches the appropriate response 
// hope this doesn't skyrocket my AWS bill D:
func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot messaage
	if message.Author.ID == session.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.Contains(message.Content, "!react4roles"):
		if message.Content[0:12] != "!react4roles" { break } //command appeared somewhere besides the front
		react4rolesCommand(session, message)
	case message.Content == "!pet" :
		session.ChannelMessageSend(message.ChannelID, "*hapy capy noises*")
	}
}

//Command: !react4roles <role name>
func react4rolesCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	// get the name of the role
	roleName := message.Content[12:]
	session.ChannelMessageSend(message.ChannelID, "React for role: " + roleName)
	fmt.Println(session.MessageReactions)
}

// TODO: handle reactions to the react4roles bot response message
func react4rolesReactionHandler(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	fmt.Println(reaction.UserID)
}

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

	// add your handlers (newMessage takes care of commands and deploys the correct response/response functions
	client.AddHandler(newMessage)
	client.AddHandler(react4rolesReactionHandler)

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
