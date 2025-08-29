package handlers

import (
	"strings"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

// called when someone reacts to the bot's message in order to get a role
func React4roleReactionAddHandler(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	// first we check to see that the reaction was made to a react4role message
	msg, err := session.ChannelMessage(reaction.ChannelID, reaction.MessageID)
	if err != nil {
		session.ChannelMessageSend(reaction.ChannelID, "Error when checking reacted message")
		return
	}

	// if the message reacted to wasn't from the bot
	if msg.Author.ID != session.State.User.ID {
		return
	}

	// we make sure that the reaction was done to a react4role message
	if !(strings.Contains(msg.Content, "React for role: ") && msg.Content[0:16] == "React for role: ") { 
		return
	}

	var roleId string

	// get the role ID from the db that corresponds to the message ID of the message that was reacted to
	row := utils.GetRoleFromMsgStatement.QueryRow(reaction.MessageID)

	err = row.Scan(&roleId)
	if err != nil {
		session.ChannelMessageSend(reaction.ChannelID, "Error when getting result from database")
		return
	}

	// add the role to the user who reacted
	err = session.GuildMemberRoleAdd(reaction.GuildID, reaction.UserID, roleId)
	if err != nil {
		session.ChannelMessageSend(reaction.ChannelID, "Error when adding role to member " + reaction.UserID)
		return
	}
}
