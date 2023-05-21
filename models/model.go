package models

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Model: app state
type Model struct {
	title     string
	eventSubs EventSubsLists
}

// New ModeL: initial model
func NewModel() Model {
	return Model{
		title:     "hello world",
		eventSubs: EventSubsLists{},
	}
}

// Init: kick off event loop
func (m Model) Init() tea.Cmd {
	return nil
}

// Update: handle  msgs
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View: return a string based on the state of the model
func (m Model) View() string {
	return m.title
}

// Cmd
type Cmd func() tea.Msg

// Msg
type Msg interface{}

// Twitch EventSubs
type EventSubsLists struct {
	Subscription []struct {
		Condition struct {
			BroadcasterUserID string `json:"broadcaster_user_id"`
		} `json:"condition"`
		Cost      int       `json:"cost"`
		CreatedAt time.Time `json:"created_at"`
		ID        string    `json:"id"`
		Status    string    `json:"status"`
		Transport struct {
			Callback string `json:"callback"`
			Method   string `json:"method"`
		} `json:"transport"`
		Type    string `json:"type"`
		Version string `json:"version"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
	Total int `json:"total"`
}
