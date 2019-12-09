package helptexts

import "github.com/Bendodroid/BendoBot/util"

var (
	Helptexts HelptextDB
)

func init() {
	// Populate the Helptext map
	util.LoadJSON("Helptexts.json", &Helptexts)
}
