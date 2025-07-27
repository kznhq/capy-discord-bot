package commands

import (
	"math/rand"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

//Command: !react4roles <role name>
func React4roleCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	// get the name of the role
	utils.RoleName = message.Content[12:] //message.Content[11] is the space character
	if utils.RoleName == "" { // (Discord shaves off trailing spaces when sending a message)
		session.ChannelMessageSend(message.ChannelID, "No role name detected. Usage: !react4role <roleName>")
	}

	msg, _ := session.ChannelMessageSend(message.ChannelID, "React for role: " + utils.RoleName)

	// parameters for the role
	color := rand.Intn(0x1000000)
	hoist := false
	permissions := int64(discordgo.PermissionViewChannel)
	mentionable := true
	params := &discordgo.RoleParams{
		Name: utils.RoleName,
		Color: &color,
		Hoist: &hoist,
		Permissions: &permissions,
		Mentionable: &mentionable,
	}

	// make the role with the above parameters
	utils.Role, _ = session.GuildRoleCreate(message.GuildID, params)
	utils.RoleAssigningMessages[msg.ID] = [2]string{utils.RoleName, string(utils.Role.ID)}
}
