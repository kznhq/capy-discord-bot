package commands

import (
	"math/rand"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func OwdCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
		session.ChannelMessageSend(message.ChannelID, "Random OW DPS: " + utils.OwDps[rand.Intn(len(utils.OwDps))])
}
