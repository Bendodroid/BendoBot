package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/errors"
	"github.com/Bendodroid/BendoBot/helptexts"
)

var err error

func PingPreRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func PingRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "Pong!")
	errors.CheckMsgSend(err, ev.GuildID, ev.ChannelID)
}

func PingPostRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func PingHelp(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, helptexts.Helptexts.Texts["ping"])
	errors.CheckMsgSend(err, ev.GuildID, ev.ChannelID)
}
