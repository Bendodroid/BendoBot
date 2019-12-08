package handlers

import (
	"strings"

	"github.com/Bendodroid/BendoBot/util"
)

func ParseInput(input string) []string {
	if !strings.HasPrefix(input, util.BotConfig.Prefix) {
		return nil
	}
	input = strings.Replace(input, util.BotConfig.Prefix, "", 1)
	return strings.Split(input, " ")
}
