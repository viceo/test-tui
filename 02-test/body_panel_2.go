package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type body_panel_2 struct{}

func newBodyPanel2() *body_panel_2 {
	return &body_panel_2{}
}

func (x body_panel_2) Init() tea.Cmd {
	return nil
}

func (x body_panel_2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return x, nil
}

func (x body_panel_2) View() string {
	return "panel 2"
}
