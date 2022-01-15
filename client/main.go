package client

import (
	"fmt"
	"strconv"
	"strings"

	"felixwie.com/producer/client/components"
	iv "felixwie.com/producer/client/pages/inventoryList"
	"felixwie.com/producer/client/styles"
	"felixwie.com/producer/logic"
	"felixwie.com/producer/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gliderlabs/ssh"
)

func GetClient(term ssh.Pty, user string) *X {
	p := &X{
		nav: components.NewSelection([]string{"Produce", "Inventory"}, styles.DefaultStyles(), func(selection string) func() tea.Msg {
			return func() tea.Msg {
				switch selection {
				case "Produce":
					return components.NavigateMsg{
						Element: components.Title{
							Content: "Produce",
						},
					}
				case "Inventory":
					return components.NavigateMsg{
						Element: iv.InventoryList{},
					}
				default:
					return nil
				}
			}
		}),
		view:         components.NewTitle("Producer"),
		style:        styles.DefaultStyles(),
		activeWindow: 0,
		term:         term,
		user:         user,
	}

	return p
}

type X struct {
	nav  tea.Model
	view tea.Model

	activeWindow int
	style        *styles.Styles

	term ssh.Pty
	user string
}

func (x X) Init() tea.Cmd {
	return nil
}

func (x X) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return x, tea.Quit
		case "tab":
			x.activeWindow = (x.activeWindow + 1) % 2
			return x, nil
		}
	case components.NavigateMsg:
		switch msg.Element.(type) {
		case components.Title:
			x.view = components.NewTitle(msg.Element.(components.Title).Content)
			x.activeWindow = 1
		case iv.InventoryList:
			x.view = iv.New()
			x.activeWindow = 1
		}
	case iv.NewEntry:
		x.view = components.NewInput([]string{"Name", "Qty"}, func(results []string) {
			intVar, _ := strconv.Atoi(results[1])
			logic.Create(models.Inventory{
				Name:              results[0],
				QuantityAvailable: intVar,
			})
		}, iv.InventoryList{})
	}

	switch x.activeWindow {
	case 0:
		ab, cmd := x.nav.Update(msg)
		x.nav = ab
		return x, cmd
	case 1:
		ab, cmd := x.view.Update(msg)
		x.view = ab
		return x, cmd
	default:
		return x, nil
	}
}

func (x X) View() string {
	var slb lipgloss.Style
	var srb lipgloss.Style
	slb = x.style.Menu
	srb = x.style.Menu

	if x.activeWindow == 0 {
		slb = slb.Copy().BorderForeground(x.style.ActiveBorderColor)
	} else {
		srb = srb.Copy().BorderForeground(x.style.ActiveBorderColor)
	}

	lb := slb.Render(x.nav.View())
	rb := srb.Render(x.view.View())

	content := lipgloss.JoinHorizontal(lipgloss.Top, lb, rb)

	content += "\n"
	content += x.createFooter()

	return x.style.App.Render(content)
}

func (x X) createFooter() string {
	w := strings.Builder{}
	h := []helpEntry{
		{"tab", "section"},
		{"↑/↓", "navigate"},
		{"enter", "select"},
		{"q", "quit"},
	}

	if _, ok := x.view.(iv.InventoryList); ok {
		h = append(h[:2], helpEntry{"n", "new entry"}, h[2])
	}

	if _, ok := x.view.(components.Input); ok {
		h = append(h[:2], helpEntry{"esc", "go back"}, h[2])
	}

	for i, v := range h {
		w.WriteString(v.Render(x.style))
		if i != len(h)-1 {
			w.WriteString(x.style.HelpDivider.String())
		}
	}

	return x.style.Footer.Copy().Render(w.String())
}

type helpEntry struct {
	key string
	val string
}

func (h helpEntry) Render(s *styles.Styles) string {
	return fmt.Sprintf("%s %s", s.HelpKey.Render(h.key), s.HelpValue.Render(h.val))
}
