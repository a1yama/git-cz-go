package components

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BreakingModel represents the breaking change confirmation component
type BreakingModel struct {
	selected bool
	cursor   int
}

// BreakingSelectedMsg is sent when the breaking change option is selected
type BreakingSelectedMsg struct {
	IsBreaking bool
}

// NewBreakingModel creates a new breaking change model
func NewBreakingModel() BreakingModel {
	return BreakingModel{
		selected: false,
		cursor:   1, // Default to "No"
	}
}

// Init initializes the breaking model
func (m BreakingModel) Init() tea.Cmd {
	return nil
}

// Update handles breaking change selection updates
func (m BreakingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
			m.cursor = 0 // Yes
		case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
			m.cursor = 1 // No
		case key.Matches(msg, key.NewBinding(key.WithKeys("y", "Y"))):
			// Quick select Yes
			return m, func() tea.Msg {
				return BreakingSelectedMsg{IsBreaking: true}
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("n", "N"))):
			// Quick select No
			return m, func() tea.Msg {
				return BreakingSelectedMsg{IsBreaking: false}
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))):
			isBreaking := m.cursor == 0
			return m, func() tea.Msg {
				return BreakingSelectedMsg{IsBreaking: isBreaking}
			}
		}
	}

	return m, nil
}

// View renders the breaking change selection
func (m BreakingModel) View() string {
	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("170")).
		Bold(true)
	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))

	options := []string{"Yes", "No"}
	result := ""

	for i, option := range options {
		cursor := "  "
		if i == m.cursor {
			cursor = "▸ "
			result += selectedStyle.Render(cursor+option) + "\n"
		} else {
			result += normalStyle.Render(cursor+option) + "\n"
		}
	}

	result += "\n" + helpStyle.Render("⚠️  Breaking changes indicate incompatible API changes (Y/N for quick select)")

	return result
}
