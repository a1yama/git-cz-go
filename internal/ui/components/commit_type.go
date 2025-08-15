package components

import (
	"fmt"
	"github.com/a1yama/git-cz-go/internal/config"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

// CommitTypeSelectedMsg is sent when a commit type is selected
type CommitTypeSelectedMsg struct {
	Type string
}

// CommitTypeModel handles the commit type selection
type CommitTypeModel struct {
	types    []config.CommitType
	useEmoji bool
	cursor   int
}

// NewCommitTypeModel creates a new commit type model
func NewCommitTypeModel(types []config.CommitType, useEmoji bool) CommitTypeModel {
	return CommitTypeModel{
		types:    types,
		useEmoji: useEmoji,
		cursor:   0,
	}
}

// Init initializes the model
func (m CommitTypeModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the model
func (m CommitTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
			if m.cursor < len(m.types)-1 {
				m.cursor++
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))):
			return m, func() tea.Msg {
				return CommitTypeSelectedMsg{Type: m.types[m.cursor].Type}
			}
		default:
			// Check if the key is a number (1-9)
			key := msg.String()
			if key >= "1" && key <= "9" {
				index := int(key[0] - '1')
				if index >= 0 && index < len(m.types) {
					return m, func() tea.Msg {
						return CommitTypeSelectedMsg{Type: m.types[index].Type}
					}
				}
			}
		}
	}

	return m, nil
}

// View renders the model
func (m CommitTypeModel) View() string {
	var result strings.Builder

	// Styles for different states
	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true)

	normalStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("39"))

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245"))

	cursorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true)

	// Render each commit type option
	for i, commitType := range m.types {
		cursor := "   "
		if i == m.cursor {
			cursor = cursorStyle.Render("▸  ")
		}

		// Build the line
		line := cursor

		// Add number
		numberStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("243"))
		line += numberStyle.Render(fmt.Sprintf("%d ", i+1))

		// Add emoji and type
		var typeText string
		if m.useEmoji && commitType.Emoji != "" {
			typeText = fmt.Sprintf("%s %s", commitType.Emoji, commitType.Type)
		} else {
			typeText = commitType.Type
		}

		if i == m.cursor {
			line += selectedStyle.Render(typeText)
		} else {
			line += normalStyle.Render(typeText)
		}

		// Add description
		line += descStyle.Render(fmt.Sprintf(" - %s", commitType.Description))

		result.WriteString(line)
		if i < len(m.types)-1 {
			result.WriteString("\n")
		}
	}

	// Add hint
	hintStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("243")).
		MarginTop(1).
		Italic(true)

	hint := hintStyle.Render("\n\n💡 Use ↑/↓ arrows or 1-9 keys for selection")

	return result.String() + hint
}
