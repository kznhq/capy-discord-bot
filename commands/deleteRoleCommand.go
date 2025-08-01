package commands

import (
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func DeleteRoleCommand (session *discordgo.Session, message *discordgo.MessageCreate) {
	roleToDelete := message.Content[12:] //content[11] is a space character
	if roleToDelete == ""{
		session.ChannelMessageSend(message.ChannelID, "Error: No role name detected. Usage is '!deleteRole <roleName>'")
	}

	for _, value := range utils.RoleAssigningMessages { // value is (role name, role id)
		if value[0] == roleToDelete {
			session.GuildRoleDelete(message.GuildID, value[1])
			session.ChannelMessageSend(message.ChannelID, "Role " + roleToDelete + " deleted")
			return
		}
	}
	session.ChannelMessageSend(message.ChannelID, "Error: couldn't find role " + roleToDelete + " to delete. Either I didn't make it or it doesn't exist")
}

