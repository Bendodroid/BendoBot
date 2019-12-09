package commands

import "github.com/bwmarrin/discordgo"

// The type for the functions every command has to implement
type CommandFunc func(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string)

type Command struct {
	PreRun  CommandFunc
	Run     CommandFunc
	PostRun CommandFunc
	Help    CommandFunc
}
