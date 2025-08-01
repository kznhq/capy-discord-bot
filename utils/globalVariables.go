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
	"!remindMe <num days>:<num hours>:<num minutes> <message is optional>": "capy will remind you after the given amount of time by replying to your message with @ turned on and repeating the inputted message if there is one",
	"!r6randomdef": "picks a random defender in Rainbow Six Siege so you don't have to say 'Guys who should I play?'",
	"!r6randomatt": "picks a random attacker in Rainbow Six Siege so you don't have to say 'Guys who should I play?'",
}

// list of all attacker operators in Rainbow Six Siege for !randomR6op command
var AttackersR6 = [38]string {
	"Rauora", "Striker", "Deimos", "Ram", "Brava", "Grim", "Sens", "Osa",
	"Flores", "Zero", "Ace", "Iana", "Kali", "Amaru", "Nokk", "Gridlock",
	"Nomad", "Maverick", "Lion", "Finka", "Dokkaebi", "Zofia", "Ying", "Jackal",
	"Hibana", "Capitao", "Blackbeard", "Buck", "Sledge", "Thatcher", "Ash", "Thermite",
	"Montagne", "Twitch", "Blitz", "IQ", "Fuze", "Glaz",
}

// list of all defender operators in Rainbow Six Siege for !randomR6op command
var DefendersR6 = [37]string {
	"Skopos", "Sentry", "Tubarao", "Fenrir", "Solis", "Azami", "Thorn", "Aruni",
	"Thunderbird", "Melusi", "Oryx", "Wamai", "Goyo", "Warden", "Mozzie", "Kaid",
	"Clash", "Maestro", "Alibi", "Vigil", "Ela", "Lesion", "Mira", "Echo",
	"Caveira", "Valkyrie", "Frost", "Mute", "Smoke", "Castle", "Pulse", "Doc",
	"Rook", "Jager", "Bandit", "Tachanka", "Kapkan",
}
