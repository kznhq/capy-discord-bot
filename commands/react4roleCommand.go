package commands

import (
	"math/rand"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

//Command: !react4roles <role name>
func React4roleCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	utils.M.Lock()

	// get the name of the role
	roleName := message.Content[12:] //message.Content[11] is the space character
	if roleName == "" { // Discord shaves off trailing spaces when sending a message
		session.ChannelMessageSend(message.ChannelID, "Error: No role name detected. Usage is '!react4role <roleName>'")
		utils.M.Unlock()
		return
	}

	msg, _ := session.ChannelMessageSend(message.ChannelID, "React for role: " + roleName)

	// we check to see if we made the role before in the server
	// if so, the entry we add will have the key of the above message so we send the message again
	// the values will be the same as the role that already exists
	for _, roleInfo := range utils.RoleAssigningMessages {
		if roleInfo[0] == roleName {
			utils.RoleAssigningMessages[msg.ID] = [3]string{roleName, roleInfo[1], roleInfo[2]}
			utils.M.Unlock()
			return
		}
	}

	// parameters for the role
	color := rand.Intn(0x1000000) //color is a hex of 2 bytes per color and 3 colors (RGB)
	hoist := false
	permissions := int64(discordgo.PermissionViewChannel)
	mentionable := true
	params := &discordgo.RoleParams{
		Name: roleName,
		Color: &color,
		Hoist: &hoist,
		Permissions: &permissions,
		Mentionable: &mentionable,
	}

	// make the role with the above parameters
	role, err := session.GuildRoleCreate(message.GuildID, params)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error: Failed to create role" + roleName)
		utils.M.Unlock()
		return
	}

	utils.RoleAssigningMessages[msg.ID] = [3]string{roleName, string(role.ID), message.GuildID}

	utils.M.Unlock()
}
