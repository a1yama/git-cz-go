package components

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BodyModel represents the body input component
type BodyModel struct {
	textarea        textarea.Model
	err             error
	lastWasNewline  bool
}

// BodySubmittedMsg is sent when the body is submitted
type BodySubmittedMsg struct {
	Body string
}

// NewBodyModel creates a new body input model
func NewBodyModel() BodyModel {
	ta := textarea.New()
	ta.Placeholder = "Provide a longer description of the change (optional - press Enter twice to continue)"
	ta.Focus()
	ta.SetWidth(80)
	ta.SetHeight(6)
	ta.ShowLineNumbers = false
	ta.CharLimit = 500

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
		case tea.KeyEnter:
			// If the textarea is empty, submit immediately
			if m.textarea.Value() == "" {
				return m, func() tea.Msg {
					return BodySubmittedMsg{Body: ""}
				}
			}
			// If last input was also Enter (double Enter), submit
			if m.lastWasNewline {
				// Remove the last newline before submitting
				value := m.textarea.Value()
				if len(value) > 0 && value[len(value)-1] == '\n' {
					value = value[:len(value)-1]
				}
				return m, func() tea.Msg {
					return BodySubmittedMsg{Body: value}
				}
			}
			m.lastWasNewline = true
			// Pass the Enter key to textarea for normal newline
			m.textarea, cmd = m.textarea.Update(msg)
			return m, cmd
		case tea.KeyCtrlD:
			// Keep Ctrl+D as alternative submit
			return m, func() tea.Msg {
				return BodySubmittedMsg{Body: m.textarea.Value()}
			}
		case tea.KeyEsc:
			// ESC is handled by the parent
			return m, nil
		default:
			// Reset the newline flag for any other key
			m.lastWasNewline = false
		}
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

// View renders the body input
func (m BodyModel) View() string {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	helpText := "📝 Body provides additional context (Press Enter twice or Ctrl+D to continue, Enter on empty to skip)"
	return m.textarea.View() + "\n\n" +
		helpStyle.Render(helpText)
}
