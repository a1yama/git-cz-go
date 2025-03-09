package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BreakingSubmittedMsg is sent when a breaking change selection is made
type BreakingSubmittedMsg struct {
	IsBreaking bool
}

// BreakingModel handles the breaking change selection
type BreakingModel struct {
	isBreaking bool
}

// NewBreakingModel creates a new breaking change model
func NewBreakingModel() BreakingModel {
	return BreakingModel{
		isBreaking: false,
	}
}

// Init initializes the model
func (m BreakingModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the model
func (m BreakingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.isBreaking = true
			return m, func() tea.Msg {
				return BreakingSubmittedMsg{IsBreaking: true}
			}
		case "n", "N":
			m.isBreaking = false
			return m, func() tea.Msg {
				return BreakingSubmittedMsg{IsBreaking: false}
			}
		case "enter":
			return m, func() tea.Msg {
				return BreakingSubmittedMsg{IsBreaking: m.isBreaking}
			}
		case "tab", "left", "right", "space":
			m.isBreaking = !m.isBreaking
		}
	}

	return m, nil
}

// View renders the model
func (m BreakingModel) View() string {
	yesStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	noStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))

	if m.isBreaking {
		yesStyle = yesStyle.Foreground(lipgloss.Color("2")).Bold(true)
	} else {
		noStyle = noStyle.Foreground(lipgloss.Color("2")).Bold(true)
	}

	view := lipgloss.JoinHorizontal(
		lipgloss.Center,
		yesStyle.Render("[Y] Yes"),
		"   ",
		noStyle.Render("[N] No"),
	)

	// Add explanation
	explanation := lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Render(
		"\n\nA breaking change means that users will need to make changes to their code when updating to your commit.")

	return view + explanation
}
