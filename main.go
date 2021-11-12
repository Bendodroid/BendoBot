package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bendodroid/BendoBot/config"
	"github.com/Bendodroid/BendoBot/errors"
	"github.com/Bendodroid/BendoBot/handlers"
	"github.com/bwmarrin/discordgo"
)

func main() {
	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + config.BotConfig.Token)
	errors.Check(err, "Error creating Discord session")

	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(handlers.MessageCreate)

	// Add Handler for ready events.
	session.AddHandler(handlers.Ready)

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()
	errors.Check(err, "Error opening connection")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = session.Close()
	errors.Check(err, "Failed to close session properly")
}
