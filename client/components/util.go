package components

import tea "github.com/charmbracelet/bubbletea"

type NavigateMsg struct {
	Element any
}

func sendNavigateMsg(element any) func() tea.Msg {
	return func() tea.Msg {
		return NavigateMsg{
			Element: element,
		}
	}
}
