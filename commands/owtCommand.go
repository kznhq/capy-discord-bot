package commands

import (
	"math/rand"

	"github.com/kznhq/capyDiscordBot/utils"

	"github.com/bwmarrin/discordgo"
)

func OwtCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
		session.ChannelMessageSend(message.ChannelID, "Random OW tank: " + utils.OwTanks[rand.Intn(len(utils.OwTanks))])
}
