package commands

import (
	"strings"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func RemindMeCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	splits := strings.Split(message.Content[10:], ":")

	if len(splits) > 3 {
		session.ChannelMessageSend(message.ChannelID, "Error: use for !remindMe is <days>:<hours>:<minutes>")
		return
	}

	// var days, hours, minutes string = "", "", ""
	_, err := strconv.Atoi(splits[0])
	if err != nil { 	// we check that we were given a proper number by trying to convert to an int
		session.ChannelMessageSend(message.ChannelID, "Error: failed to parse number of days from input, are you sure it's an integer?")
			return
	}
	if len(splits) > 1 {
		_, err := strconv.Atoi(splits[1])
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error: failed to parse number of hours from input, are you sure it's an integer?")
			return
		}
	}
	if len(splits) > 2 {
		_, err := strconv.Atoi(splits[2])
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error: failed to parse number of minutes from input, are you sure it's an integer?")
			return
		}
	}
	go remind(session, message, splits[0], splits[1], splits[2])
}

// helper function called as goroutine that actually handles the reminding
func remind(session *discordgo.Session, message *discordgo.MessageCreate, days string, hours string, minutes string) {
	intDays, err := strconv.Atoi(days)
	if err != nil { // I don't think these errors should trigger because we checked earlier, but Go wants me to
		session.ChannelMessageSend(message.ChannelID, "Error: couldn't parse days from given input in timer goroutine")
	}
	intHours, err := strconv.Atoi(hours)
	if err != nil { 
		session.ChannelMessageSend(message.ChannelID, "Error: couldn't parse hours from given input in timer goroutine")
	}

	sleepHours := intDays * 24 + intHours // time.ParseDuration() doesn't parse days :(

	duration, err := time.ParseDuration(strconv.Itoa(sleepHours) + "h" + minutes + "m")
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "Error: couldn't parse time from given input in timer goroutine")
		return
	}

	// capy sends an acknowledgement so the user knows the reminder worked
	session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend {
		Content: "Ok, I'll remind you in " + days + " days, " + hours + " hours, and " + minutes + " minutes.",
		Reference: &discordgo.MessageReference {
			MessageID: message.ID,
			ChannelID: message.ChannelID,
			GuildID: message.GuildID,
		},
	})
	
	time.Sleep(time.Minute * time.Duration(duration.Minutes()))

	session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend {
		Content: "Time's up!",
		Reference: &discordgo.MessageReference {
			MessageID: message.ID,
			ChannelID: message.ChannelID,
			GuildID: message.GuildID,
		},
	})
}
