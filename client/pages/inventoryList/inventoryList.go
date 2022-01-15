package inventorylist

import (
	"fmt"
	"io"
	"strings"

	"felixwie.com/producer/logic"
	"felixwie.com/producer/models"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type listItem struct {
}

func (i listItem) FilterValue() string                       { return "" }
func (i listItem) Height() int                               { return 1 }
func (i listItem) Spacing() int                              { return 0 }
func (i listItem) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (i listItem) Render(w io.Writer, m list.Model, index int, data models.Inventory) {
	str := fmt.Sprintf("%d. %s", index+1, data.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

type InventoryList struct {
	list   list.Model
	data   *[]models.Inventory
	choice models.Inventory
}

func New() InventoryList {
	d, _ := logic.GetAll[models.Inventory](&logic.QueryOptions{})

	l := list.New(*d, listItem{}, 20, 14)
	l.Title = "YALLO"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return InventoryList{
		data: d,
		list: l,
	}
}

func (i InventoryList) Init() tea.Cmd {
	return nil
}

type NewEntry struct{}

func (i InventoryList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		i.list.SetWidth(msg.Width)
		return i, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "n":
			cmds = append(cmds, i.sendNewMessage)
		case "enter":
			if _, ok := i.list.SelectedItem().(models.Inventory); ok {
				i.choice = i.list.SelectedItem().(models.Inventory)
			}

			return i, nil
		}
	}

	return i, tea.Batch(cmds...)
}

func (i InventoryList) View() string {
	s := strings.Builder{}
	s.WriteString("Inventory:\n")

	s.WriteString("\n")

	for i, v := range *i.data {
		s.WriteString(fmt.Sprintf("%d. %s\n", i, v.Name))
	}

	return s.String()
}

func (i InventoryList) sendNewMessage() tea.Msg {
	return NewEntry{}
}
