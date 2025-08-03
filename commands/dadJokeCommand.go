package commands

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"

	"github.com/bwmarrin/discordgo"
)

type DadJoke struct {
	Id string `json:id`
	Joke string `json:joke`
	Status int `json:status`
}

func DadJokeCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil { 
		session.ChannelMessageSend(message.ChannelID, "Error when creating request for dad joke D:")
	}

	// icanhazdadjoke API docs kindly requested making a custom User Agent header so they can monitor usage
	req.Header.Set("User-Agent", "thanks for doing this! (https://github.com/kznhq/capy-discord-bot)")
	// icanhazdadjoke requires specifying how you want your joke formatted
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error when doing GET request for dad joke D:")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error code when getting dad joke: %d", resp.StatusCode))
		return
	}

	var dadJoke DadJoke
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error when reading response")
		return
	}

	err = json.Unmarshal(body, &dadJoke)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error when unmarshaling response")
	}

	session.ChannelMessageSend(message.ChannelID, dadJoke.Joke)
}
