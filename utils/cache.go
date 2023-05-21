package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/goproslowyo/eventsub/models"
)

type Cache struct {
	CacheDir      string
	CacheFilePath string
	EventSubs     models.EventSubsLists
}

func (c Cache) GetCache() Cache {
	// Check for cache data in $XDG_CACHE_HOME or $HOME/.cache.
	// If not found, download and cache
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Warnf("failed to get user cache directory: %w", err)
	}
	cacheFilePath := filepath.Join(cacheDir, "twitch-eventsub/subscriptions.json")
	cacheFile, err := os.Open(cacheFilePath)
	if err != nil {
		log.Infof("failed to open cache file: %s", err)
	}
	if _, err := os.Stat(cacheFilePath); err == nil {
		// Cache file exists, no need to download
		// Parse json and return EventSubsLists
		cacheFile, err := os.Open(cacheFilePath)
		if err != nil {
			log.Infof("failed to open cache file: %s", err)
		}
		defer cacheFile.Close()

		var eventSubsLists models.EventSubsLists
		json.NewDecoder(cacheFile).Decode(&eventSubsLists)
		if err != nil {
			log.Errorf("failed to respose: %s", err)
		}

		return Cache{
			CacheDir:      cacheDir,
			CacheFilePath: cacheFilePath,
			EventSubs:     eventSubsLists,
		}
	} else {
		// Cache file does not exist
		// download and cache by calling the Twitch CLI
		resp, err := RunTwitchCli([]string{"api", "get", "eventsub/subscriptions"})
		if err != nil {
			log.Errorf("failed to download subscriptions: %s", err)
		}
		eventSubs := models.EventSubsLists{}
		json.Unmarshal(resp, &eventSubs)
		cacheFile.Write(resp)
		defer cacheFile.Close()
		return Cache{
			CacheDir:      cacheDir,
			CacheFilePath: cacheFilePath,
			EventSubs:     eventSubs,
		}
	}
}
