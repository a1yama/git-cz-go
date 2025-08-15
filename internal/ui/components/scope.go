package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ScopeModel represents the scope input component
type ScopeModel struct {
	textInput textinput.Model
	err       error
}

// ScopeSubmittedMsg is sent when the scope is submitted
type ScopeSubmittedMsg struct {
	Scope string
}

// NewScopeModel creates a new scope input model
func NewScopeModel() ScopeModel {
	ti := textinput.New()
	ti.Placeholder = "e.g., api, ui, core (optional - press Enter to skip)"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 60

	return ScopeModel{
		textInput: ti,
	}
}

// Init initializes the scope model
func (m ScopeModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles scope input updates
func (m ScopeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Allow empty scope (optional field)
			return m, func() tea.Msg {
				return ScopeSubmittedMsg{Scope: m.textInput.Value()}
			}
		case tea.KeyEsc:
			// ESC is handled by the parent
			return m, nil
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View renders the scope input
func (m ScopeModel) View() string {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	
	return m.textInput.View() + "\n\n" +
		helpStyle.Render("📦 Scope narrows the context of the change (optional)")
}