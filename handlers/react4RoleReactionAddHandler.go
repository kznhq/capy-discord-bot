package handlers

import (
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

// called when someone reacts to the bot's message in order to get a role
func React4roleReactionAddHandler(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	var roleId string

	// get the role ID from the db that corresponds to the message ID of the message that was reacted to
	row := utils.GetRoleFromMsgStatement.QueryRow(reaction.MessageID)

	err := row.Scan(&roleId)
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
