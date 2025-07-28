package commands

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"

	"github.com/bwmarrin/discordgo"
)

type Fact struct {
	Id string `json:id`
	Text string `json:text`
}

func FactCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	resp, err := http.Get("https://uselessfacts.jsph.pl/random.json")
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error when doing GET request for fact D:")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error code when getting fact: %d", resp.StatusCode))
		return
	}

	var fact Fact
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error when reading response")
		return
	}

	err = json.Unmarshal(body, &fact)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error when unmarshaling response")
	}

	session.ChannelMessageSend(message.ChannelID, fact.Text)
}
