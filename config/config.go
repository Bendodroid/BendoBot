package config

import "github.com/Bendodroid/BendoBot/util"

var (
	BotConfig Config
)

func init() {
	// Load the Bot config
	util.LoadJSON("Config.json", &BotConfig)
}
