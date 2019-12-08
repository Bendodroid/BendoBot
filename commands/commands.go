package commands

import (
	"github.com/bwmarrin/discordgo"
)

type CommandFunc func(s *discordgo.Session, ev *discordgo.MessageCreate, msg *[]string)

type Command struct {
	PreRun  CommandFunc
	Run     CommandFunc
	PostRun CommandFunc
	Help    CommandFunc
}

func NewCommand(preRun CommandFunc, run CommandFunc, postRun CommandFunc, help CommandFunc) Command {
	return Command{PreRun: preRun, Run: run, PostRun: postRun, Help: help}
}
