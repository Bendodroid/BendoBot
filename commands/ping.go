package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/errors"
)

var err error

func PingPreRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func PingRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "Pong!")
	errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
}

func PingPostRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func PingHelp(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "```text\nping\n\tIf the BendoBot works, it replies with 'Pong!'```")
	errors.CheckMsgSendErr(err, ev.GuildID, ev.ChannelID)
}
