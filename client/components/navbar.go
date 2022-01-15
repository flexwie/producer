package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	selected map[int]struct{}
	output   string
}

func InitModel() model {
	return model{
		choices:  []string{"Inventory", "Produce"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// quit program
		case "ctrl+c", "q":
			return m, tea.Quit
		case "1":
			m.output = "Inventory"
		case "2":
			m.output = "Produce"
		}
	}

	return m, nil
}

func (m model) View() string {
	baseStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7D56F4"))

	var s string
	s += baseStyle.Copy().Bold(true).PaddingRight(3).Render("PRODUCER")

	for i, choice := range m.choices {
		content := fmt.Sprintf("[%d] %s", i+1, choice)
		s += baseStyle.Copy().PaddingRight(1).Render(content)
	}

	s += fmt.Sprintf("\n\nChoice is: %s\n", m.output)

	return s
}
