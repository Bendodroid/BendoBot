package errors

import (
	"fmt"
	"log"
)

// checkErr panics with an error message if err != nil
func Check(err error, logMsg string) {
	if err != nil {
		log.Panicf("[ERROR] %s", logMsg)
	}
}

// checkMsgSendErr is a wrapper for checkErr with a msg prefilled for ChannelMessageSend errors
func CheckMsgSendErr(err error, gid string, chanid string) {
	Check(err, fmt.Sprintf("Failed sending message in guild: %s in channel: %s", gid, chanid))
}
