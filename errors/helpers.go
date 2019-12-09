package errors

import (
	"fmt"
	"log"
)

// Check panics with an error message if err != nil
func Check(err error, logMsg string) {
	if err != nil {
		log.Panicf("[ERROR] %s", logMsg)
	}
}

// CheckMsgSend is a wrapper for errors.Check with a msg prefilled for ChannelMessageSend errors
func CheckMsgSend(err error, gid string, chid string) {
	Check(err, fmt.Sprintf("Failed sending message in guild: %s in channel: %s", gid, chid))
}
