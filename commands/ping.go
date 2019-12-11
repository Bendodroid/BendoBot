package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/errors"
	"github.com/Bendodroid/BendoBot/helptexts"
)

var err error

func pingPreRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func pingRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "Pong!")
	errors.CheckMsgSend(err, ev.GuildID, ev.ChannelID)
}

func pingPostRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func pingHelp(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, helptexts.Helptexts.Texts["ping"])
	errors.CheckMsgSend(err, ev.GuildID, ev.ChannelID)
}
