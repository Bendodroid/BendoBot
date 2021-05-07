package config

import "github.com/Bendodroid/BendoBot/util"

var (
	// BotConfig contains the currently active configuration options
	BotConfig Config
)

func init() {
	// Load the Bot config
	util.LoadJSON("config.json", &BotConfig)
}
