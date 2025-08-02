package handlers

import (
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

// called when someone reacts to the bot's message in order to get a role
func React4roleReactionAddHandler(session *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	if roleInfo, ok := utils.RoleAssigningMessages[reaction.MessageID]; ok { //if there is a reaction to one of the messages that the bot sent for assigning a role
		session.GuildMemberRoleAdd(reaction.GuildID, reaction.UserID, roleInfo[1]) // 0 is name, 1 is role ID, 2 is guild ID
	}
}
