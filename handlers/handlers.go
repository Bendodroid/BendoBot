package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/commands"
	"github.com/Bendodroid/BendoBot/errors"
)

// MessageCreate is called every time a new message is created in a channel that the bot has access to
func MessageCreate(s *discordgo.Session, ev *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if ev.Author.ID == s.State.User.ID {
		return
	}
	// Parse the message content and split into arguments
	input := ParseInput(ev.Content)
	// If there are no arguments, it was no command
	if input == nil {
		return
	}

	command := commands.CommandMap[strings.ToLower(input[0])]
	// Check if command is nil before doing sth with it
	if command.Run == nil {
		return
	}
	// Check for help command
	if len(input) > 1 && input[1] == "help" {
		command.Help(s, ev, &input)
		return
	}
	// Call all the methods for the associated command
	command.PreRun(s, ev, &input)
	command.Run(s, ev, &input)
	command.PostRun(s, ev, &input)
}

// Ready is called when receiving the "ready" event
func Ready(s *discordgo.Session, _ *discordgo.Ready) {
	// Set the playing status
	err := s.UpdateGameStatus(0, "with gophers...")
	errors.Check(err, "Failed setting custom status")
}
