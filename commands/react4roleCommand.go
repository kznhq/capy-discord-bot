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
	}

	msg, _ := session.ChannelMessageSend(message.ChannelID, "React for role: " + roleName)

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
