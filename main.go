package main

import (
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/charmbracelet/log"
	m "github.com/goproslowyo/eventsub/models"
)

func main() {
	log.Info("starting app")
	log.Warn("warning app")
	log.Debug("debugging app")
	log.Error("erroring app")
	// New program with initial model and options
	p := tea.NewProgram(m.NewModel(), tea.WithAltScreen())
	// Run
	_, err := p.Run()
	if err != nil {
		log.Fatalf("could not run program: %s", err)
	}
}
