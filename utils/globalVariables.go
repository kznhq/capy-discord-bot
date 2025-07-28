package utils

import (
	"github.com/bwmarrin/discordgo"
)

// stores the global variables used by different files

// map that stores the messages that are used to assign roles (the messages created by !react4role)
//the key is the message ID (which are strings), the value is the (role name, role ID) that message assigns
var RoleAssigningMessages = make(map[string][2]string) 

// the role being created in one call of !react4role
var Role *discordgo.Role

// the name of the role being created in one call of !react4role
var RoleName string

// map of all the commands, used for the !help command to show all bot functionalities
var CommandMap = map[string]string {
	"!pet": "do it",
	"!react4role <role name>": "create a role with the given name, react to the bot's message to get it or un-react to no longer be part of it",
	"!deleteRole <role name>": "delete the role with the given name, only works for roles made by capy",
	"!fact": "capy tells you a random fun fact. It pulls these from some APIs so I can't guarantee they're actually true",
}
