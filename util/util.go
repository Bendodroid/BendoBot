package util

import (
	"github.com/Bendodroid/BendoBot/commands"
	"github.com/Bendodroid/BendoBot/config"
)

var (
	BotConfig  config.Config
	CommandMap = make(map[string]commands.Command)
)

func init() {
	BotConfig = config.Load("Config.json")

	CommandMap["ping"] = commands.NewCommand(
		commands.PingPreRun, commands.PingRun, commands.PingPostRun, commands.PingHelp)
}
