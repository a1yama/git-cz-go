package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BodySubmittedMsg is sent when a body is submitted
type BodySubmittedMsg struct {
	Body string
}

// BodyModel handles the commit body input
type BodyModel struct {
	textarea      textarea.Model
	maxLineLength int
}

// NewBodyModel creates a new body model
func NewBodyModel(maxLineLength int) BodyModel {
	ta := textarea.New()
	ta.Placeholder = "Provide a longer description of the change (optional)"
	ta.Focus()
	ta.SetWidth(60)
	ta.SetHeight(10)
	ta.ShowLineNumbers = false

	return BodyModel{
		textarea:      ta,
		maxLineLength: maxLineLength,
	}
}

// Init initializes the model
func (m BodyModel) Init() tea.Cmd {
	return textarea.Blink
}

// Update handles updates for the model
func (m BodyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Ctrl+Enter または Enter を処理
		if msg.Type == tea.KeyCtrlJ || msg.Type == tea.KeyEnter && msg.Alt {
			return m, func() tea.Msg {
				return BodySubmittedMsg{Body: m.textarea.Value()}
			}
		}

		// ESC キーの処理
		if msg.Type == tea.KeyEsc {
			if !m.textarea.Focused() {
				return m, func() tea.Msg {
					return BodySubmittedMsg{Body: m.textarea.Value()}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

// View renders the model
func (m BodyModel) View() string {
	// Main textarea view
	view := m.textarea.View()

	// Check for lines that exceed max length
	var longLines []int
	lines := strings.Split(m.textarea.Value(), "\n")
	for i, line := range lines {
		if len(line) > m.maxLineLength {
			longLines = append(longLines, i+1)
		}
	}

	// Add warning about long lines if any
	if len(longLines) > 0 {
		warning := lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(
			fmt.Sprintf("Warning: Lines %v exceed the maximum length of %d characters",
				longLines, m.maxLineLength))
		view += "\n\n" + warning
	}

	// Add help text
	view += "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Render("Tips:") +
		"\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("- Include motivation for the change") +
		"\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("- Contrast with previous behavior") +
		"\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("- Press Ctrl+Enter to submit")

	return view
}
