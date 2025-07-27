package handlers

import (
	"strings"

	"github.com/kznhq/capyDiscordBot/commands"

	"github.com/bwmarrin/discordgo"
)

// reads all messages and dispatches the appropriate response 
// hope this doesn't skyrocket my AWS bill D:
func NewMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot message
	if message.Author.ID == session.State.User.ID {
		return
	}

	// Respond to messages
	switch {

		case strings.Contains(message.Content, "!react4role"):
			if message.Content[0:11] != "!react4role" { break } //command appeared somewhere besides the front
			commands.React4roleCommand(session, message)

		case message.Content == "!pet" :
			session.ChannelMessageSend(message.ChannelID, "*hapy capy noises*")

		case strings.Contains(message.Content, "!deleteRole"):
			if message.Content[0:11] != "!deleteRole" { break } //command appeared somewhere besides the front
			commands.DeleteRole(session, message, message.Content)

	}
}

