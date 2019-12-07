package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/src/errors"
)

var (
	err error
)

// messageCreate is called every time a new message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, ev *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if ev.Author.ID == s.State.User.ID {
		return
	}

	if strings.ToLower(ev.Content) == "ping" {
		_, err = s.ChannelMessageSend(ev.ChannelID, "Pong!")
		errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
	}

	if strings.ToLower(ev.Content) == "pong" {
		_, err = s.ChannelMessageSend(ev.ChannelID, "Ping!")
		errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
	}
}

// ready is called when the bot receives the "ready" event from Discord.
func Ready(s *discordgo.Session, _ *discordgo.Ready) {
	// Set the playing status.
	err := s.UpdateStatus(0, "with gophers...")
	errors.CheckErr(err, "Failed setting custom status")
}
