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
		session.ChannelMessageSend(message.ChannelID, "Error: use is '!remindMe <days>:<hours>:<minutes>'")
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
		_, err := strconv.Atoi(strings.Split(splits[2], " ")[0])
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, "Error: failed to parse number of minutes from input, are you sure it's an integer?")
			return
		}
	}

	// initiate go routine for timer
	go remind(session, message)
}

// helper function called as goroutine that actually handles the reminding
func remind(session *discordgo.Session, message *discordgo.MessageCreate) {
	splits := strings.Split(message.Content[10:], ":")
	days := splits[0]
	hours := splits[1]
	minutes := ""
	msg := ""
	if len(strings.Split(splits[2], " ")) > 1 { // if there's a message, separate the minutes from the message
		minutes = strings.Split(splits[2], " ")[0]
		msg = splits[2][strings.Index(splits[2], " ") + 1:] // there is a space character after the minutes, everything after that is the message
	} else {	// if no message, take from the second colon to the end (days:hours:minutes)
		minutes = splits[2]
	}

	intDays, err := strconv.Atoi(days)
	if err != nil { 
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
		Content: "Time's up! " + msg,
		Reference: &discordgo.MessageReference {
			MessageID: message.ID,
			ChannelID: message.ChannelID,
			GuildID: message.GuildID,
		},
	})
}
