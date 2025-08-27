package handlers

import (
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

// called when someone removes their reaction to the bot's role-assigning message in order to get off of a role
func React4roleReactionRemoveHandler(session *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	var roleId string

	// get the role ID from the db that corresponds to the message ID of the message that was unreacted to
	row := utils.GetRoleFromMsgStatement.QueryRow(reaction.MessageID)

	err := row.Scan(&roleId)
	if err != nil {
		session.ChannelMessageSend(reaction.ChannelID, "Error when getting result from database")
		return
	}

	// remove the role from that user
	err = session.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, roleId)
	if err != nil {
		session.ChannelMessageSend(reaction.ChannelID, "Error when removing role from member " + reaction.UserID)
		return
	}
	// if roleInfo, ok := utils.RoleAssigningMessages[reaction.MessageID]; ok { //if there is a reaction to one of the messages that the bot sent for assigning a role
	// 	session.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, roleInfo[1]) // 0 is name, 1 is role ID, 2 is guild ID
	// }
}
