package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/errors"
	"github.com/Bendodroid/BendoBot/util"
)

// messageCreate is called every time a new message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, ev *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if ev.Author.ID == s.State.User.ID {
		return
	}

	command := ParseInput(ev.Content)
	if command == nil {
		return
	}

	if len(command) > 1 && command[1] == "help" {
		util.CommandMap[strings.ToLower(command[0])].Help(s, ev, &command)
		return
	}

	util.CommandMap[strings.ToLower(command[0])].PreRun(s, ev, &command)
	util.CommandMap[strings.ToLower(command[0])].Run(s, ev, &command)
	util.CommandMap[strings.ToLower(command[0])].PostRun(s, ev, &command)
}

// ready is called when the bot receives the "ready" event from Discord.
func Ready(s *discordgo.Session, _ *discordgo.Ready) {
	// Set the playing status.
	err := s.UpdateStatus(0, "with gophers...")
	errors.Check(err, "Failed setting custom status")
}
