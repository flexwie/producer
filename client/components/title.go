package components

import tea "github.com/charmbracelet/bubbletea"

type Title struct {
	Content string
}

func NewTitle(content string) Title {
	return Title{
		Content: content,
	}
}

func (t Title) Init() tea.Cmd {
	return nil
}

func (t Title) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t Title) View() string {
	return t.Content
}
