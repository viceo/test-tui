package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	zone "github.com/lrstanley/bubblezone"
)

var (
	normal = lipgloss.NewStyle().
		Padding(0, 4).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#444")).
		Foreground(lipgloss.Color("#888"))

	hover = lipgloss.NewStyle().
		Padding(0, 4).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#50FA7B")).
		Background(lipgloss.Color("#50FA7B")).
		BorderBackground(lipgloss.Color("#50FA7B")).
		Foreground(lipgloss.Color("#000"))
)

type button struct {
	id      string
	hovered bool
}

// Define Tea msgs
type SwitchToPanel1 struct{}
type SwitchToPanel2 struct{}

func (x button) Init() tea.Cmd {
	return nil
}

func (x button) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.MouseMsg:
		// Click detection
		if msg.Action == tea.MouseActionMotion {
			x.hovered = zone.Get(x.id).InBounds(msg) == true
		}
		if msg.Action == tea.MouseAction(tea.MouseButtonLeft) {
			if zone.Get(x.id).InBounds(msg) {
				switch x.id {
				case "btn-1":
					cmd = func() tea.Msg { return SwitchToPanel1{} }
				case "btn-2":
					cmd = func() tea.Msg { return SwitchToPanel2{} }
				}
			}
		}
	}

	return x, cmd
}

func (x button) View() string {
	btnStyle := normal
	if x.hovered {
		btnStyle = hover
	}

	return zone.Mark(x.id, btnStyle.Render(fmt.Sprintf(" %s ", x.id)))
}
