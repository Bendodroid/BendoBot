package helptexts

import "github.com/Bendodroid/BendoBot/util"

var (
	// Helptexts is a struct containing a map for all helptexts
	Helptexts HelptextDB
)

func init() {
	// Populate the Helptext map
	util.LoadJSON("Helptexts.json", &Helptexts)
}
