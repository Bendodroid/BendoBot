package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bendodroid/BendoBot/src/config"
	"github.com/Bendodroid/BendoBot/src/errors"
	"github.com/Bendodroid/BendoBot/src/handlers"

	"github.com/bwmarrin/discordgo"
)

// Global variables for later
var (
	CONFIG config.Config
)

// init reads config and so on
func init() {
	config.LoadConfig("Config.json", &CONFIG)
}

func main() {
	// Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + CONFIG.Token)
	errors.CheckErr(err, "Error creating Discord session")

	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(handlers.MessageCreate)

	// Add Handler for ready events.
	session.AddHandler(handlers.Ready)

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()
	errors.CheckErr(err, "Error opening connection")

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = session.Close()
	errors.CheckErr(err, "Failed to close session properly")
}
