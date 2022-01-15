package styles

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	ActiveBorderColor   lipgloss.Color
	InacticeBorderColor lipgloss.Color

	App    lipgloss.Style
	Header lipgloss.Style

	Menu             lipgloss.Style
	MenuCursor       lipgloss.Style
	MenuItem         lipgloss.Style
	SelectedMenuItem lipgloss.Style

	Title    lipgloss.Style
	TitleBox lipgloss.Style

	Footer      lipgloss.Style
	HelpKey     lipgloss.Style
	HelpValue   lipgloss.Style
	HelpDivider lipgloss.Style

	Error      lipgloss.Style
	ErrorTitle lipgloss.Style
	ErrorBody  lipgloss.Style
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.ActiveBorderColor = lipgloss.Color("62")
	s.InacticeBorderColor = lipgloss.Color("236")

	s.App = lipgloss.NewStyle().Margin(1, 2)

	s.Header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("62")).
		Align(lipgloss.Right).
		Bold(true)

	s.Menu = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(s.InacticeBorderColor).
		Padding(1, 2).
		MarginRight(1).
		Width(24)

	s.MenuCursor = lipgloss.NewStyle().
		Foreground(lipgloss.Color("207")).
		SetString("> ")

	s.MenuItem = lipgloss.NewStyle().
		PaddingLeft(2)

	s.Title = lipgloss.NewStyle().
		Padding(0, 2)

	s.TitleBox = lipgloss.NewStyle().
		BorderStyle(lipgloss.Border{
			Top:         "─",
			Bottom:      "─",
			Left:        "│",
			Right:       "│",
			TopLeft:     "╭",
			TopRight:    "┬",
			BottomLeft:  "├",
			BottomRight: "┴",
		})

	s.Footer = lipgloss.NewStyle().
		MarginTop(1)

	s.HelpKey = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	s.HelpValue = lipgloss.NewStyle().
		Foreground(lipgloss.Color("239"))

	s.HelpDivider = lipgloss.NewStyle().
		Foreground(lipgloss.Color("237")).
		SetString(" • ")

	s.Error = lipgloss.NewStyle().
		Padding(1)

	s.ErrorTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("230")).
		Background(lipgloss.Color("204")).
		Bold(true).
		Padding(0, 1)

	s.ErrorBody = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252")).
		MarginLeft(2).
		Width(52)

	return s
}
