package handlers

import (
	"strings"

	"github.com/kznhq/capyDiscordBot/commands"
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

// reads all messages and dispatches the appropriate response 
func NewMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot message
	if message.Author.ID == session.State.User.ID {
		return
	}

	// Respond to messages
	switch {
		// check that the command was somewhere in the message AND the command is at the beginning of the message.
		// I'm pretty sure trailing spaces are cut off before sending in Discord but here it ensures that 
		// the user put something after the command that they want the role to be called instead of just
		// "!react4role". Same logic for the others that use this
		case strings.Contains(message.Content, "!react4role ") && message.Content[0:12] == "!react4role ":
			commands.React4roleCommand(session, message)

		case message.Content == "!pet" :
			session.ChannelMessageSend(message.ChannelID, "*hapy capy noises*")

		case strings.Contains(message.Content, "!deleteRole ") && message.Content[0:12] == "!deleteRole ":
			commands.DeleteRoleCommand(session, message)

		case message.Content == "!help":
			msg := ""
			for k, v := range utils.CommandMap {
				msg = msg + "\n**" + k + "**: " + v
			}
			session.ChannelMessageSend(message.ChannelID, msg)

		case message.Content == "!fact":
			commands.FactCommand(session, message)	

		case strings.Contains(message.Content, "!remindMe ") && message.Content[0:10] == "!remindMe ":
			commands.RemindMeCommand(session, message)

		case message.Content == "!rrod":
			commands.R6RandomDefCommand(session, message)	

		case message.Content == "!rroa":
			commands.R6RandomAttCommand(session, message)	

		case message.Content == "!owt":
			commands.OwtCommand(session, message)	

		case message.Content == "!owd":
			commands.OwdCommand(session, message)	

		case message.Content == "!ows":
			commands.OwsCommand(session, message)	
	}
}
