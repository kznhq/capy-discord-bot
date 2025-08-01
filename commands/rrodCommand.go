package commands

import (
	"math/rand"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func R6RandomDefCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
		session.ChannelMessageSend(message.ChannelID, "Random R6 defender: " + utils.R6Defenders[rand.Intn(len(utils.R6Defenders))])
}
