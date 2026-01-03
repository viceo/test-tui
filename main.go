package main

import (
	"fmt"
	"os"

	"github.com/76creates/stickers/flexbox"
	"github.com/charmbracelet/bubbles/textinput"
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
		Foreground(lipgloss.Color("#000"))

	clicked = lipgloss.NewStyle().
		Padding(0, 4).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0000")).
		Background(lipgloss.Color("#FF0000")).
		Foreground(lipgloss.Color("#000"))
)

type app_model struct {
	input    textinput.Model
	button_1 tea.Model
	button_2 tea.Model
	width    int
	height   int
}
type button_model struct {
	zone_id string
	clicked bool
	hovered bool
	fn_msg  func()
}

func (m app_model) Init() tea.Cmd {
	return nil
}

func (m button_model) Init() tea.Cmd {
	return nil
}

func (m app_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String():
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.button_1, cmd = m.button_1.Update(msg)
	cmds = append(cmds, cmd)
	m.button_2, cmd = m.button_2.Update(msg)
	cmds = append(cmds, cmd)
	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m button_model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		// Click detection
		if msg.Action == tea.MouseActionMotion {
			m.hovered = zone.Get(m.zone_id).InBounds(msg) == true
		}
		if msg.Action == tea.MouseAction(tea.MouseButtonLeft) {
			if zone.Get(m.zone_id).InBounds(msg) {
				m.fn_msg()
				m.clicked = true
			}
		}
	}
	return m, nil
}

func (m app_model) View() string {
	btn1 := m.button_1.View()
	btn2 := m.button_2.View()
	inpt := m.input.View()

	view := flexbox.NewHorizontal(m.width, m.height)
	columns := []*flexbox.Column{
		view.NewColumn().AddCells(
			flexbox.NewCell(1, 1).
				SetStyle(lipgloss.NewStyle().Background(lipgloss.Color("#FF0000"))).
				SetContent(lipgloss.NewStyle().Margin(1).Render(inpt)),
		),
		view.NewColumn().AddCells(
			flexbox.NewCell(1, 1).SetContent(lipgloss.JoinHorizontal(lipgloss.Left, btn1, btn2)),
		),
	}
	view.AddColumns(columns)
	return zone.Scan(view.Render())
}

func (m button_model) View() string {
	btnStyle := normal
	// Hover check: scan updates internal state, then check if mouse is over zone
	if m.hovered {
		btnStyle = hover
	}

	if m.clicked {
		btnStyle = clicked
	}

	button := zone.Mark(m.zone_id, btnStyle.Render(" Click Me! "))

	return button
}

func main() {
	// Initialize the global zone manager
	zone.NewGlobal()

	// Input "bubble"
	input := textinput.New()
	input.Placeholder = "URL"
	input.Focus()
	input.Cursor.Blink = true

	p := tea.NewProgram(
		app_model{
			input: input,
			button_1: &button_model{
				zone_id: "btn-1",
				fn_msg: func() {
					// fmt.Println("click btn-1")
				},
			},
			button_2: &button_model{
				zone_id: "btn-2",
				fn_msg: func() {
					// fmt.Println("click btn-2")
				},
			},
		},
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(), // Required for hover!
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
