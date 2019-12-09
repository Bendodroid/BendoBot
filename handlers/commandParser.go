package handlers

import (
	"strings"

	"github.com/Bendodroid/BendoBot/config"
)

// ParseInput checks if the prefix is present, if yes, it returns a list of arguments
func ParseInput(input string) []string {
	if !strings.HasPrefix(input, config.BotConfig.Prefix) {
		return nil
	}
	input = strings.Replace(input, config.BotConfig.Prefix, "", 1)
	return strings.Split(input, " ")
}
