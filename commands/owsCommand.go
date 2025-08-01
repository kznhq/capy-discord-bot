package commands

import (
	"math/rand"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func OwsCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
		session.ChannelMessageSend(message.ChannelID, "Random OW support: " + utils.OwSupports[rand.Intn(len(utils.OwSupports))])
}
