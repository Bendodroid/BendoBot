package commands

import (
	"github.com/bwmarrin/discordgo"

	"github.com/Bendodroid/BendoBot/errors"
	"github.com/Bendodroid/BendoBot/helptexts"
)

// Error to be used later
var err error

func init() {
	helptexts.DB["ping"] = "```text\nping\n\tIf the BendoBot works, it replies with 'Pong!'```"
}

func pingPreRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func pingRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, "Pong!")
	errors.CheckMsgSend(err, ev.GuildID, ev.ChannelID)
}

func pingPostRun(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
}

func pingHelp(s *discordgo.Session, ev *discordgo.MessageCreate, args *[]string) {
	_, err = s.ChannelMessageSend(ev.ChannelID, helptexts.DB["ping"])
	errors.CheckMsgSend(err, ev.GuildID, ev.ChannelID)
}
