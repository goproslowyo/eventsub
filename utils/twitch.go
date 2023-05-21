package utils

import (
	"os"
	"os/exec"

	log "github.com/charmbracelet/log"
)

// Check for the binary (/usr/local/bin/twitch)
func checkBinary() string {
	var twitchCli string
	var err error
	if os.Getenv("EVENTSUB_TWITCH_CLI_PATH") != "" {
		twitchCli = os.Getenv("EVENTSUB_TWITCH_CLI_PATH")
	} else {
		twitchCli, err = exec.LookPath("twitch")
	}

	// Display bubbletea error modal
	if err != nil {
		log.Fatal("Twitch doesn't appear to be installed!\n")
		os.Exit(1)
	}

	log.Debugf("Twitch CLI '%s' found in PATH'\n", twitchCli)
	return twitchCli
}
