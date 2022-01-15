package components

import (
	"strings"

	"felixwie.com/producer/client/styles"
	tea "github.com/charmbracelet/bubbletea"
)

type Selection struct {
	Items        []string
	SelectedItem int
	styles       *styles.Styles

	OnSelection func(selection string) func() tea.Msg
}

func NewSelection(items []string, style *styles.Styles, onSelection func(selection string) func() tea.Msg) *Selection {
	return &Selection{
		Items:       items,
		styles:      style,
		OnSelection: onSelection,
	}
}

func (s Selection) Init() tea.Cmd {
	return nil
}

func (s Selection) View() string {
	w := strings.Builder{}

	for i, item := range s.Items {
		if i == s.SelectedItem {
			w.WriteString(s.styles.MenuCursor.String())
			w.WriteString(s.styles.SelectedMenuItem.Render(item))
		} else {
			w.WriteString(s.styles.MenuItem.Render(item))
		}

		if i < len(s.Items)-1 {
			w.WriteRune('\n')
		}
	}

	return w.String()
}

func (s Selection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "k", "up", "1":
			if s.SelectedItem > 0 {
				s.SelectedItem--
				cmds = append(cmds, s.sendActiveMessage)
			}
		case "j", "down", "2":
			if s.SelectedItem < len(s.Items)-1 {
				s.SelectedItem++
				cmds = append(cmds, s.sendActiveMessage)
			}
		case "enter":
			cmds = append(cmds, s.OnSelection(s.Items[s.SelectedItem]))
		}
	}

	return s, tea.Batch(cmds...)
}

type ActiveMsg struct {
	Name  string
	Index int
}

func (s *Selection) sendActiveMessage() tea.Msg {
	if s.SelectedItem >= 0 && s.SelectedItem < len(s.Items) {
		return ActiveMsg{
			Name:  s.Items[s.SelectedItem],
			Index: s.SelectedItem,
		}
	}

	return nil
}
