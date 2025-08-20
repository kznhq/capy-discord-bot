package commands

import (
	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func DeleteRoleCommand (session *discordgo.Session, message *discordgo.MessageCreate) {
	utils.M.Lock()

	roleToDelete := message.Content[12:] //content[11] is a space character
	if roleToDelete == ""{
		session.ChannelMessageSend(message.ChannelID, "Error: No role name detected. Usage is '!deleteRole <roleName>'")
		utils.M.Unlock()
		return
	}

	for _, value := range utils.RoleAssigningMessages {
		if value[0] == roleToDelete && value[2] == message.GuildID {	// if the role name and guild ID match
			if session.GuildRoleDelete(message.GuildID, value[1]) != nil {
				session.ChannelMessageSend(message.ChannelID, "Error: Failed to delete role " + roleToDelete)
				utils.M.Unlock()
				return
			}
			session.ChannelMessageSend(message.ChannelID, "Role " + roleToDelete + " deleted")
			utils.M.Unlock()
			return
		}
	}
	session.ChannelMessageSend(message.ChannelID, "Error: couldn't find role " + roleToDelete + " to delete. Either I didn't make it or it doesn't exist")

	utils.M.Unlock()
}

