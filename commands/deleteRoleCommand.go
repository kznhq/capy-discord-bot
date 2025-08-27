package commands

import (
	"fmt"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func DeleteRoleCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	roleToDelete := message.Content[12:] //message.Content[11] is a space character

	if roleToDelete == ""{
		session.ChannelMessageSend(message.ChannelID, "Error: No role name detected. Usage is '!deleteRole <roleName>'")
		return
	}

	// first we query the db to find the role ID
	var roleId string
	row := utils.GetRoleStatement.QueryRow(roleToDelete, message.GuildID)
	err := row.Scan(&roleId)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error: couldn't find role " + roleToDelete + " to delete. Either I didn't make it or it doesn't exist")
		return
	}

	// delete the role from the server
	if session.GuildRoleDelete(message.GuildID, roleId) != nil {
		session.ChannelMessageSend(message.ChannelID, "Error: Failed to delete role " + roleToDelete + " from server. Try again")
		return // if we don't delete from the server we don't want to delete from the database because that'd be mismatched records
	}

	// delete all rows with the matching role and guild ID from database
	_, err = utils.RoleDb.Exec("DELETE FROM roleassigningmessagestable WHERE rolename = ? AND guildid = ?", roleToDelete, message.GuildID)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error when deleting role " + roleToDelete + " from database");
		fmt.Println(err)
		return
	}

	session.ChannelMessageSend(message.ChannelID, "Role " + roleToDelete + " deleted")
}

