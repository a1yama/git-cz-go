package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// ScopeSubmittedMsg is sent when a scope is submitted
type ScopeSubmittedMsg struct {
	Scope string
}

// ScopeModel handles the scope input
type ScopeModel struct {
	textInput       textinput.Model
	suggestions     []string
	selected        int
	showSuggestions bool
}

// NewScopeModel creates a new scope model
func NewScopeModel() ScopeModel {
	ti := textinput.New()
	ti.Placeholder = "scope (optional)"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	return ScopeModel{
		textInput:       ti,
		suggestions:     []string{},
		selected:        -1,
		showSuggestions: false,
	}
}

// SetSuggestions sets the scope suggestions
func (m *ScopeModel) SetSuggestions(suggestions []string) {
	m.suggestions = suggestions
}

// Init initializes the model
func (m ScopeModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles updates for the model
func (m ScopeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// If a suggestion is selected, use it
			if m.showSuggestions && m.selected >= 0 && m.selected < len(m.filteredSuggestions()) {
				scope := m.filteredSuggestions()[m.selected]
				return m, func() tea.Msg {
					return ScopeSubmittedMsg{Scope: scope}
				}
			}
			// Otherwise use the text input
			return m, func() tea.Msg {
				return ScopeSubmittedMsg{Scope: m.textInput.Value()}
			}
		case "down":
			if m.showSuggestions {
				m.selected = min(m.selected+1, len(m.filteredSuggestions())-1)
			}
		case "up":
			if m.showSuggestions {
				m.selected = max(m.selected-1, 0)
			}
		case "tab":
			// Toggle suggestions visibility
			m.showSuggestions = !m.showSuggestions
			if m.showSuggestions && m.selected == -1 && len(m.filteredSuggestions()) > 0 {
				m.selected = 0
			}
		}
	}

	// Only show suggestions if we have some input and there are matching suggestions
	if len(m.textInput.Value()) > 0 && len(m.filteredSuggestions()) > 0 {
		m.showSuggestions = true
	} else {
		m.showSuggestions = false
		m.selected = -1
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View renders the model
func (m ScopeModel) View() string {
	view := m.textInput.View()

	// Add suggestions if showing
	if m.showSuggestions {
		suggestions := m.filteredSuggestions()
		if len(suggestions) > 0 {
			view += "\n\nSuggestions (tab to focus, ↑/↓ to navigate):\n"
			for i, s := range suggestions {
				if i == m.selected {
					view += lipgloss.NewStyle().Background(lipgloss.Color("4")).Foreground(lipgloss.Color("15")).Render(" > "+s+" ") + "\n"
				} else {
					view += "   " + s + "\n"
				}
			}
		}
	}

	return view
}

// filteredSuggestions returns suggestions that match the current input
func (m ScopeModel) filteredSuggestions() []string {
	if m.textInput.Value() == "" {
		return m.suggestions
	}

	var filtered []string
	for _, s := range m.suggestions {
		// Simple contains check - could be improved with more sophisticated matching
		if strings.Contains(strings.ToLower(s), strings.ToLower(m.textInput.Value())) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

// Helper functions for min and max
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
