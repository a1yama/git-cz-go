package components

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BodyModel represents the body input component
type BodyModel struct {
	textarea textarea.Model
	err      error
}

// BodySubmittedMsg is sent when the body is submitted
type BodySubmittedMsg struct {
	Body string
}

// NewBodyModel creates a new body input model
func NewBodyModel() BodyModel {
	ta := textarea.New()
	ta.Placeholder = "Provide a longer description of the change (optional - press Ctrl+D to continue)"
	ta.Focus()
	ta.SetWidth(80)
	ta.SetHeight(6)
	ta.ShowLineNumbers = false

	return BodyModel{
		textarea: ta,
	}
}

// Init initializes the body model
func (m BodyModel) Init() tea.Cmd {
	return textarea.Blink
}

// Update handles body input updates
func (m BodyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlD:
			// Submit the body
			return m, func() tea.Msg {
				return BodySubmittedMsg{Body: m.textarea.Value()}
			}
		case tea.KeyEsc:
			// ESC is handled by the parent
			return m, nil
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

// View renders the body input
func (m BodyModel) View() string {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	return m.textarea.View() + "\n\n" +
		helpStyle.Render("📝 Body provides additional context (Ctrl+D to continue, optional)")
}
