package bubble

import (
	//"fmt"
	//"os"

	"log"
	"projet/src/common"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var mess string
var playerInput string

type nameInput struct {
	textInput textinput.Model
	err       error
	quitting  bool
}

func initInput() nameInput {
	ti := textinput.New()
	ti.Placeholder = "Fever"
	ti.Focus()
	ti.CharLimit = 16

	return nameInput{textInput: ti}
}

func (m nameInput) Init() tea.Cmd {
	return textinput.Blink
}

func (m nameInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		common.PlaySFX("typing")
		switch msg.String() {
		case "enter", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m nameInput) View() tea.View {
	var c *tea.Cursor
	if !m.textInput.VirtualCursor() {
		c = m.textInput.Cursor()
		c.Y += lipgloss.Height(m.headerView(mess))
	}

	str := lipgloss.JoinVertical(lipgloss.Top, m.headerView(mess), m.textInput.View())
	if m.quitting {
		str += "\n"
	}

	v := tea.NewView(str)
	v.Cursor = c
	return v
}

func StartInput(s string) string {
	mess = s
	playerInput = ""

	p := tea.NewProgram(initInput())
	finalModel, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	if m, ok := finalModel.(nameInput); ok {
		playerInput = m.textInput.Value()
	}
	return playerInput
}

func (m nameInput) headerView(mess string) string { return mess + "\n" }
