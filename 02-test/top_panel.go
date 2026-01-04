package main

import (
	"github.com/76creates/stickers/flexbox"
	tea "github.com/charmbracelet/bubbletea"
)

type top_panel struct {
	flex_container *flexbox.HorizontalFlexBox
	btn_1          tea.Model
	btn_2          tea.Model
}

func (x top_panel) Init() tea.Cmd {
	return nil
}

func (x top_panel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	x.btn_1, cmd = x.btn_1.Update(msg)
	cmds = append(cmds, cmd)
	x.btn_2, cmd = x.btn_2.Update(msg)
	cmds = append(cmds, cmd)
	return x, tea.Batch(cmds...)
}

func (x top_panel) View() string {
	x.flex_container.SetColumns([]*flexbox.Column{
		x.flex_container.NewColumn().AddCells(
			flexbox.NewCell(1, 1).
				SetContent(x.btn_1.View()),
		),
		x.flex_container.NewColumn().AddCells(
			flexbox.NewCell(1, 1).
				SetContent(x.btn_2.View()),
		),
	})

	return x.flex_container.Render()
}
