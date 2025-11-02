package handlers

import (
	"strings"
	"math/rand"

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
			for _, k := range utils.CommandNames {
				msg = msg + "\n**" + k + "**: " + utils.CommandMap[k]
			}
			session.ChannelMessageSend(message.ChannelID, msg)

		case message.Content == "!fact":
			commands.FactCommand(session, message)	

		case message.Content == "!dadJoke":
			commands.DadJokeCommand(session, message)	

		case strings.Contains(message.Content, "!remindMe ") && message.Content[0:10] == "!remindMe ":
			commands.RemindMeCommand(session, message)

		case message.Content == "!rd":
			commands.R6RandomDefCommand(session, message)	

		case message.Content == "!ra":
			commands.R6RandomAttCommand(session, message)	

		case message.Content == "!owt":
			commands.OwtCommand(session, message)	

		case message.Content == "!owd":
			commands.OwdCommand(session, message)	

		case message.Content == "!ows":
			commands.OwsCommand(session, message)	

		case strings.Contains(strings.ToLower(message.Content), "i'm ") && strings.ToLower(message.Content[0:4]) == "i'm ": // classic dad joke right here
			num := rand.Intn(7)
			if num == 1 {	// make it a random chance that capy replies so it doesn't get too annoying
				session.ChannelMessageSend(message.ChannelID, "Hi" + message.Content[3:] + ", I'm capy!")
			}

		case strings.Contains(message.Content, "I am ") && message.Content[0:5] == "I am ":
			num := rand.Intn(7)
			if num == 1 {
				session.ChannelMessageSend(message.ChannelID, "Hi" + message.Content[4:] + ", I'm capy!")
			}
	}
}
