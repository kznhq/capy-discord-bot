package handlers

import (
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func React4roleReactionRemoveHandler(session *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	if _, ok := utils.RoleAssigningMessages[reaction.MessageID]; ok { //if there is a reaction to one of the messages that the bot sent for assigning a role
		session.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, utils.RoleName)
	}
	session.GuildMemberRoleRemove(reaction.GuildID, reaction.UserID, utils.Role.ID)
}
