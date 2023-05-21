package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

// Check for cache data in $XDG_CACHE_HOME or $HOME/.cache. If not found, download and cache
func fetchCache() error {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return fmt.Errorf("failed to get user cache directory: %w", err)
	}

	cacheFilePath := filepath.Join(cacheDir, "twitch-eventsub/subscriptions.json")

	// Check if cache file exists
	if _, err := os.Stat(cacheFilePath); err == nil {
		// Cache file exists, no need to download
		return nil
	}

	// Cache file does not exist, download by calling the Twitch CLI
	// resp, err :=
	if err != nil {
		log.Errorf("failed to download cache file: %s", err)
	}

	cacheFile, err := os.Create(cacheFilePath)
	if err != nil {
		log.Errorf("failed to create cache file: %s", err)
	}
	defer cacheFile.Close()

	return nil
}
