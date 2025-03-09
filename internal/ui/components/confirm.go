package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ConfirmMsg is sent when the user confirms or cancels
type ConfirmMsg struct {
	Confirmed bool
}

// ConfirmModel handles the confirmation of the commit message
type ConfirmModel struct {
	confirmed bool
}

// NewConfirmModel creates a new confirm model
func NewConfirmModel() ConfirmModel {
	return ConfirmModel{
		confirmed: true, // Default to confirmed
	}
}

// Init initializes the model
func (m ConfirmModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the model
func (m ConfirmModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.confirmed = true
			return m, func() tea.Msg {
				return ConfirmMsg{Confirmed: true}
			}
		case "n", "N":
			m.confirmed = false
			return m, func() tea.Msg {
				return ConfirmMsg{Confirmed: false}
			}
		case "enter":
			return m, func() tea.Msg {
				return ConfirmMsg{Confirmed: m.confirmed}
			}
		case " ", "space":
			// スペースキーのトグル処理を明示的に追加
			m.confirmed = !m.confirmed
			return m, nil
		}

		// 他のキーマッチングも確認
		if msg.Type == tea.KeySpace {
			m.confirmed = !m.confirmed
			return m, nil
		}
	}

	return m, nil
}

// View renders the model
func (m ConfirmModel) View() string {
	confirmStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	cancelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))

	if m.confirmed {
		confirmStyle = confirmStyle.Foreground(lipgloss.Color("2")).Bold(true)
	} else {
		cancelStyle = cancelStyle.Foreground(lipgloss.Color("9")).Bold(true)
	}

	view := lipgloss.JoinHorizontal(
		lipgloss.Center,
		confirmStyle.Render("[Y] Commit"),
		"   ",
		cancelStyle.Render("[N] Cancel"),
	)

	// Add help text
	help := "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render(
		"Press Enter to confirm your choice")

	return view + help
}
