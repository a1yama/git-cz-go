package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SubjectSubmittedMsg is sent when a subject is submitted
type SubjectSubmittedMsg struct {
	Subject string
}

// SubjectModel handles the commit subject input
type SubjectModel struct {
	textInput  textinput.Model
	maxLength  int
	validInput bool
}

// NewSubjectModel creates a new subject model
func NewSubjectModel(maxLength int) SubjectModel {
	ti := textinput.New()
	ti.Placeholder = "Write a concise description of the change"
	ti.Focus()
	ti.CharLimit = maxLength
	ti.Width = 50

	return SubjectModel{
		textInput:  ti,
		maxLength:  maxLength,
		validInput: false,
	}
}

// Init initializes the model
func (m SubjectModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles updates for the model
func (m SubjectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(m.textInput.Value()) > 0 {
				return m, func() tea.Msg {
					return SubjectSubmittedMsg{Subject: m.textInput.Value()}
				}
			}
		}
	}

	// Update validation status
	m.validInput = len(m.textInput.Value()) > 0

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View renders the model
func (m SubjectModel) View() string {
	// Calculate remaining characters
	currentLength := len(m.textInput.Value())
	remaining := m.maxLength - currentLength

	// Style for the character counter
	counterStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	if remaining < 10 {
		counterStyle = counterStyle.Foreground(lipgloss.Color("9")) // Red for low remaining
	}

	// Main text input view
	view := m.textInput.View()

	// Add character counter and validation hint
	view += "\n\n" + counterStyle.Render(fmt.Sprintf("%d/%d characters", currentLength, m.maxLength))

	// Add validation message if needed
	if !m.validInput && currentLength > 0 {
		view += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render("Subject cannot be empty")
	}

	// Add helper text
	view += "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Render("Tips:") +
		"\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("- Use imperative, present tense: \"add\" not \"added\"") +
		"\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("- Don't capitalize the first letter") +
		"\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("- No period at the end")

	return view
}
