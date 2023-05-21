package utils

import (
	"os"
	"os/exec"

	log "github.com/charmbracelet/log"
)

// Check for the binary (/usr/local/bin/twitch)
func GetTwitchCliPath() string {
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
	// Check for ~/.config/twitch-cli/.twitch-cli.env and return warning if not found
	// and hint to run `twitch login`
	if _, err := os.Stat("~/.config/twitch-cli/.twitch-cli.env"); err != nil {
		log.Warn("Twitch CLI config not found!\n")
		log.Warn("Run `twitch login` to login to Twitch\n")
	}

	log.Debugf("Twitch CLI '%s' found in PATH'\n", twitchCli)
	log.Debug("Twitch CLI seems configured!\n")
	return twitchCli
}

func RunTwitchCli(args []string) ([]byte, error) {
	twitchCli := GetTwitchCliPath()
	cmd := exec.Command(twitchCli, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("failed to run twitch cli: %s", err)
		return nil, err
	}
	return out, nil
}
