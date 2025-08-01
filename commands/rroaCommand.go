package commands

import (
	"math/rand"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func R6RandomAttCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
		session.ChannelMessageSend(message.ChannelID, "Random R6 attacker: " + utils.R6Attackers[rand.Intn(len(utils.R6Attackers))])
}
