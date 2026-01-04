package main

import (
	"github.com/76creates/stickers/flexbox"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

type app struct {
	flex_container *flexbox.FlexBox
	top_panel      tea.Model
	body_panel     tea.Model
}

func (x app) Init() tea.Cmd {
	return nil
}

func (x *app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		x.flex_container.SetWidth(msg.Width)
		x.flex_container.SetHeight(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyCtrlC.String():
			return x, tea.Quit
		}
	case SwitchToPanel1:
		x.body_panel = newBodyPanel1()
	case SwitchToPanel2:
		x.body_panel = newBodyPanel2()
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	x.top_panel, cmd = x.top_panel.Update(msg)
	cmds = append(cmds, cmd)

	x.body_panel, cmd = x.body_panel.Update(msg)
	cmds = append(cmds, cmd)

	return x, tea.Batch(cmds...)
}

func (x app) View() string {
	x.flex_container.SetRows([]*flexbox.Row{
		x.flex_container.NewRow().AddCells(
			flexbox.NewCell(12, 1).
				SetContent(zone.Mark("top-panel", x.top_panel.View())),
		),
		x.flex_container.NewRow().AddCells(
			flexbox.NewCell(12, 12).
				SetContent(zone.Mark("body-panel", x.body_panel.View())),
		),
	})
	return zone.Scan(x.flex_container.Render())
}

func main() {
	zone.NewGlobal()
	p := tea.NewProgram(
		&app{
			flex_container: flexbox.New(0, 0),
			top_panel: &top_panel{
				flex_container: flexbox.NewHorizontal(0, 0),
				btn_1: &button{
					id: "btn-1",
				},
				btn_2: &button{
					id: "btn-2",
				},
			},
			body_panel: newBodyPanel1(),
		},
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
