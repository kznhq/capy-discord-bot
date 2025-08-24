package commands

import (
	"math/rand"
	"fmt"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func React4roleCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	// get the name of the role
	roleName := message.Content[12:] //message.Content[11] is the space character
	if roleName == "" { // Discord shaves off trailing spaces when sending a message
		session.ChannelMessageSend(message.ChannelID, "Error: No role name detected. Usage is '!react4role <roleName>'")
		return
	}

	newRole := false // boolean we'll check later to see whether we need to create a role or not

	// first we check to see if we have made this role in this server before
	rows, err := utils.GetRoleStatement.Query(roleName, message.GuildID)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error: failed to check database to see if this role already exists")
		return
	}

	if rows.Next() { // some result came when searching for this role name in this guild which means that we made it before so we won't want to create a new role
		newRole = false
	} else { // this is if no rows were found
		// if the role name exists in the guild but not because the bot made it, we don't make a new one
		existingRoles, err := session.GuildRoles(message.GuildID)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error: Failed to check server for existing roles")
			return
		}
		for _, role := range existingRoles {
			if role.Name == roleName {
				session.ChannelMessageSend(message.ChannelID, "Error: Role already exists in server. I didn't make it but someone else did.")
				return
			}
		}

		// at this point, we didn't make the role before and no mods made it in this guild already so we will want to create a new role
		newRole = true
	}

	var role *discordgo.Role

	// if this role didn't exist before anywhere, we create one
	if newRole {
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
		role, err = session.GuildRoleCreate(message.GuildID, params)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error: Failed to create role" + roleName)
			return
		}
	}

	// we send the message before adding to db because this message's ID is the primary key
	// we do this instead of the message that called the command because we need
	// this message ID in order to check that when a reaction is added or removed
	// to a message that the ID of that message matches one of these, if so then we
	// add or remove someone to a role. Also we don't have to worry about someone
	// deleting their message calling this command since capy won't delete its own
	// message for this
	msg, _ := session.ChannelMessageSend(message.ChannelID, "React for role: " + roleName)

	// add the role to the db
	_, err = utils.RoleDb.Exec("INSERT INTO roleassigningmessagestable (messageid, rolename, roleid, guildid) VALUES (?, ?, ?, ?)", msg.ID, role.Name, role.ID, message.GuildID)
	if err != nil {
		fmt.Println(err)
		return
	}
}
