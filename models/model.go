package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// You generally won't need this unless you're processing stuff with
// complicated ANSI escape sequences. Turn it on if you notice flickering.
//
// Also keep in mind that high performance rendering only works for programs
// that use the full size of the terminal. We're enabling that below with
// tea.EnterAltScreen().
const useHighPerformanceRenderer = false

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Model: app state
type Model struct {
	eventSubs     EventSubsLists
	ready         bool
	title         string
	TwitchCliPath string
	viewport      viewport.Model
}

// New ModeL: initial model
func NewModel() Model {
	return Model{
		eventSubs:     EventSubsLists{},
		title:         "hello skyfire & world",
		TwitchCliPath: "",
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

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.TwitchCliPath)
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}
	return m, nil
}

// https://github.com/charmbracelet/bubbletea/blob/master/examples/pager/main.go
// https://camo.githubusercontent.com/9a8aa88b9c0f2a48a5622e18c8c911b5efcf885c19e4c29d7a25088ec8fe6826/68747470733a2f2f73747566662e636861726d2e73682f627562626c65732d6578616d706c65732f76696577706f72742e676966
func (m Model) headerView() string {
	title := titleStyle.Render("Mr. Pager")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

// View: return a string based on the state of the model
func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
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

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	width       = 96
	columnWidth = 30
)

func DisplayErrorModal(str string) {
	subtle := lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
	buttonStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginTop(1)
	activeButtonStyle := buttonStyle.Copy().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#F25D94")).
		MarginRight(2).
		Underline(true)
	acceptErrorButton := activeButtonStyle.Render("ERROR")

	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(str)
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, acceptErrorButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, buttons)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	doc := strings.Builder{}
	doc.WriteString(dialog + "\n\n")
}
