package main

import (
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/charmbracelet/log"
	"github.com/goproslowyo/eventsub/models"
	"github.com/goproslowyo/eventsub/utils"
)

func main() {
	log.Info("starting app")
	log.Warn("warning app")
	log.Debug("debugging app")
	log.Error("erroring app")
	// New program with initial model and options
	m := models.NewModel()
	m.TwitchCliPath = utils.GetTwitchCliPath()
	p := tea.NewProgram(m, tea.WithAltScreen())
	// Run
	_, err := p.Run()
	if err != nil {
		log.Fatalf("could not run program: %s\n blame skyfire", err)
	}
}
