package utils

import (
	"github.com/bwmarrin/discordgo"
)

var RoleAssigningMessages = make(map[string][2]string) //the key is the message ID (which are strings), the value is the (role name, role ID) that message assigns
var Role *discordgo.Role
var RoleName string

