package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/errors"
)

var err error

func PingPreRun(s *discordgo.Session, ev *discordgo.MessageCreate, msg *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "This is the PreRun")
	errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
}

func PingRun(s *discordgo.Session, ev *discordgo.MessageCreate, msg *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "This is the Run")
	errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
}

func PingPostRun(s *discordgo.Session, ev *discordgo.MessageCreate, msg *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "This is the PostRun")
	errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
}

func PingHelp(s *discordgo.Session, ev *discordgo.MessageCreate, msg *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "This is the Helptext")
	errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
}
