package components

import (
	"fmt"
	"strings"

	iv "felixwie.com/producer/client/pages/inventoryList"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type Input struct {
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode

	cb       func([]string)
	returnTo any
}

func NewInput(questions []string, cb func([]string), returnTo any) Input {
	i := Input{
		inputs:   make([]textinput.Model, len(questions)),
		cb:       cb,
		returnTo: returnTo,
	}

	var t textinput.Model
	for idx, v := range questions {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 32

		t.Placeholder = v

		if idx == 0 {
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		i.inputs[idx] = t
	}

	return i
}

func (i Input) Init() tea.Cmd {
	return textinput.Blink
}

func (i Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return i, sendNavigateMsg(i.returnTo)
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()
			cmds := make([]tea.Cmd, 0)

			if s == "enter" && i.focusIndex == len(i.inputs) {
				results := make([]string, len(i.inputs))

				for idx, v := range i.inputs {
					results[idx] = v.Value()
				}

				i.cb(results)
				cmds = append(cmds, sendNavigateMsg(iv.InventoryList{}))
				return i, tea.Batch(cmds...)
			}

			if s == "up" || s == "shift+tab" {
				i.focusIndex--
			} else {
				i.focusIndex++
			}

			if i.focusIndex > len(i.inputs) {
				i.focusIndex = 0
			} else if i.focusIndex < 0 {
				i.focusIndex = len(i.inputs)
			}

			cmds = make([]tea.Cmd, len(i.inputs))
			for idx := 0; idx <= len(i.inputs)-1; idx++ {
				if idx == i.focusIndex {
					// Set focused state
					cmds[idx] = i.inputs[idx].Focus()
					i.inputs[idx].PromptStyle = focusedStyle
					i.inputs[idx].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				i.inputs[idx].Blur()
				i.inputs[idx].PromptStyle = noStyle
				i.inputs[idx].TextStyle = noStyle
			}

			return i, tea.Batch(cmds...)
		}
	}

	cmd := i.updateInputs(msg)

	return i, cmd
}

func (i Input) View() string {
	var b strings.Builder

	for idx := range i.inputs {
		b.WriteString(i.inputs[idx].View())
		if idx < len(i.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if i.focusIndex == len(i.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}

func (i Input) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(i.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for idx := range i.inputs {
		i.inputs[idx], cmds[idx] = i.inputs[idx].Update(msg)
	}

	return tea.Batch(cmds...)
}
