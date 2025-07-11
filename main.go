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
	case strings.Contains(message.Content, "!react4role"):
		if message.Content[0:11] != "!react4role" { break } //command appeared somewhere besides the front
		react4roleCommand(session, message)
	case message.Content == "!pet" :
		session.ChannelMessageSend(message.ChannelID, "*hapy capy noises*")
	case strings.Contains(message.Content, "!deleteRole"):
		if message.Content[0:11] != "!deleteRole" { break } //command appeared somewhere besides the front
		deleteRole(session, message, message.Content)
	}
}

func deleteRole (session *discordgo.Session, message *discordgo.MessageCreate, content string) {
	roleToDelete := content[12:] //content[11] is a space character
	if roleToDelete == ""{
		session.ChannelMessageSend(message.ChannelID, "Error: No role name detected. Usage: !deleteRole <roleName>")
	}

	for _, value := range roleAssigningMessages { // value is (role name, role id)
		if value[0] == roleToDelete {
			session.GuildRoleDelete(message.GuildID, value[1])
			session.ChannelMessageSend(message.ChannelID, "Role " + roleToDelete + " deleted")
			return
		}
	}
	session.ChannelMessageSend(message.ChannelID, "Error: couldn't find role " + roleToDelete + " to delete. Either I didn't make it or it doesn't exist")
}

//Command: !react4roles <role name>
func react4roleCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	// get the name of the role
	roleName = message.Content[12:] //message.Content[11] is the space character
	if roleName == "" { // (Discord shaves off trailing spaces when sending a message)
		session.ChannelMessageSend(message.ChannelID, "No role name detected. Usage: !react4role <roleName>")
	}

	msg, _ := session.ChannelMessageSend(message.ChannelID, "React for role: " + roleName)
	color := 0x208470
	hoist := false
	permissions := int64(discordgo.PermissionViewChannel)
	mentionable := true
	params := &discordgo.RoleParams{
		Name: roleName,
		Color: &color,
		Hoist: &hoist,
		Permissions: &permissions,
		Mentionable: &mentionable,
	}
	role, _ = session.GuildRoleCreate(message.GuildID, params)
	roleAssigningMessages[msg.ID] = [2]string{roleName, string(role.ID)}
}

func react4roleReactionAddHandler(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	if _, ok := roleAssigningMessages[reaction.MessageID]; ok { //if there is a reaction to one of the messages that the bot sent for assigning a role
		session.ChannelMessageSend(reaction.ChannelID, "Reaction found")
		session.GuildMemberRoleAdd(reaction.GuildID, reaction.UserID, role.ID)
	}
}

func react4roleReactionRemoveHandler(session *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	if _, ok := roleAssigningMessages[reaction.MessageID]; ok { //if there is a reaction to one of the messages that the bot sent for assigning a role
		session.ChannelMessageSend(reaction.ChannelID, "Reaction removed")
		session.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, roleName)
	}
	session.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, role.ID)
}

var (
	roleAssigningMessages = make(map[string][2]string) //the key is the message ID (which are strings), the value is the (role name, role ID) that message assigns
	role *discordgo.Role
	roleName string
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

	// add your handlers (newMessage takes care of commands and deploys the correct response/response functions
	client.AddHandler(newMessage)
	client.AddHandler(react4roleReactionAddHandler)
	client.AddHandler(react4roleReactionRemoveHandler)

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
