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

var CommandNames = [10]string {
	"!pet",
	"!react4role <role name>",
	"!deleteRole <role name>",
	"!fact",
	"!remindMe <num days>:<num hours>:<num minutes> <message is optional>",
	"!rroa",
	"!rrod",
	"!owt",
	"!ows",
	"!owd",
}

// map of all the commands, used for the !help command to show all bot functionalities
var CommandMap = map[string]string {
	"!pet": "do it",
	"!react4role <role name>": "create a role with the given name, react to the bot's message to get it or un-react to no longer be part of it",
	"!deleteRole <role name>": "delete the role with the given name, only works for roles made by capy",
	"!fact": "capy tells you a random fun fact. It pulls these from some APIs so I can't guarantee they're actually true",
	"!remindMe <num days>:<num hours>:<num minutes> <message is optional>": "capy will remind you after the given amount of time by replying to your message with @ turned on and repeating the inputted message if there is one",
	"!rrod": "picks a random defender from Rainbow Six Siege so you don't have to say 'Guys who should I play?', command is named so it's easy to fast to type on mobile :)",
	"!rroa": "picks a random attacker from Rainbow Six Siege",
	"!owt": "picks a random tank from Overwatch",
	"!ows": "picks a random support from Overwatch",
	"!owd": "picks a random DPS from Overwatch",
}

// list of all attacker operators in Rainbow Six Siege for !rroa command
var R6Attackers = [38]string {
	"Rauora", "Striker", "Deimos", "Ram", "Brava", "Grim", "Sens", "Osa",
	"Flores", "Zero", "Ace", "Iana", "Kali", "Amaru", "Nokk", "Gridlock",
	"Nomad", "Maverick", "Lion", "Finka", "Dokkaebi", "Zofia", "Ying", "Jackal",
	"Hibana", "Capitao", "Blackbeard", "Buck", "Sledge", "Thatcher", "Ash", "Thermite",
	"Montagne", "Twitch", "Blitz", "IQ", "Fuze", "Glaz",
}

// list of all defender operators in Rainbow Six Siege for !rrod command
var R6Defenders = [37]string {
	"Skopos", "Sentry", "Tubarao", "Fenrir", "Solis", "Azami", "Thorn", "Aruni",
	"Thunderbird", "Melusi", "Oryx", "Wamai", "Goyo", "Warden", "Mozzie", "Kaid",
	"Clash", "Maestro", "Alibi", "Vigil", "Ela", "Lesion", "Mira", "Echo",
	"Caveira", "Valkyrie", "Frost", "Mute", "Smoke", "Castle", "Pulse", "Doc",
	"Rook", "Jager", "Bandit", "Tachanka", "Kapkan",
}

// list of all tanks in Overwatch for !owt
var OwTanks = [13]string {
	"D.VA", "Doomfist", "Hazard", "Junker Queen", "Mauga", "Orisa", "Ramattra", "Reinhardt",
	"Roadhog", "Sigma", "Winston", "Wrecking Ball", "Zarya",
}

// list of all DPS in Overwatch for !owd
var OwDps = [19]string {
	"Ashe", "Bastion", "Cassidy", "Echo", "Freja", "Genji", "Hanzo", "Junkrat",
	"Mei", "Pharah", "Reaper", "Sojourn", "Soldier 76", "Somba", "Symmetra", "Torbjorn",
	"Tracer", "Venture", "Widowmaker",
}

// list of all supports in Overwatch for !ows
var OwSupports = [11]string {
	"Ana", "Baptiste", "Brigitte", "Illari", "Juno", "Kiriko", "Lifeweaver", "Lucio",
	"Mercy", "Moira", "Zenyatta",
}
