package components

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// FooterModel represents the footer input component
type FooterModel struct {
	typeInput  textinput.Model
	valueInput textinput.Model
	stage      int // 0: type, 1: value
}

// FooterSubmittedMsg is sent when the footer is submitted
type FooterSubmittedMsg struct {
	FooterType  string
	FooterValue string
}

// NewFooterModel creates a new footer input model
func NewFooterModel() FooterModel {
	ti1 := textinput.New()
	ti1.Placeholder = "e.g., Closes, Fixes, BREAKING CHANGE (optional - press Enter to skip)"
	ti1.Focus()
	ti1.CharLimit = 50
	ti1.Width = 60

	ti2 := textinput.New()
	ti2.Placeholder = "e.g., #123, issue description"
	ti2.CharLimit = 200
	ti2.Width = 60

	return FooterModel{
		typeInput:  ti1,
		valueInput: ti2,
		stage:      0,
	}
}

// Init initializes the footer model
func (m FooterModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles footer input updates
func (m FooterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.stage == 0 {
				// If no type is entered, skip footer entirely
				if m.typeInput.Value() == "" {
					return m, func() tea.Msg {
						return FooterSubmittedMsg{
							FooterType:  "",
							FooterValue: "",
						}
					}
				}
				// Move to value input
				m.stage = 1
				m.valueInput.Focus()
				return m, textinput.Blink
			} else {
				// Submit footer
				return m, func() tea.Msg {
					return FooterSubmittedMsg{
						FooterType:  strings.TrimSpace(m.typeInput.Value()),
						FooterValue: strings.TrimSpace(m.valueInput.Value()),
					}
				}
			}
		case tea.KeyEsc:
			if m.stage == 1 {
				// Go back to type input
				m.stage = 0
				m.typeInput.Focus()
				m.valueInput.Blur()
				return m, textinput.Blink
			}
			// ESC in stage 0 is handled by parent
			return m, nil
		}
	}

	if m.stage == 0 {
		m.typeInput, cmd = m.typeInput.Update(msg)
	} else {
		m.valueInput, cmd = m.valueInput.Update(msg)
	}

	return m, cmd
}

// View renders the footer input
func (m FooterModel) View() string {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("86"))

	if m.stage == 0 {
		return labelStyle.Render("Footer Type:") + "\n" +
			m.typeInput.View() + "\n\n" +
			helpStyle.Render("🔗 Footer references issues or breaking changes (optional)")
	}

	return labelStyle.Render("Footer Type: ") + m.typeInput.Value() + "\n\n" +
		labelStyle.Render("Footer Value:") + "\n" +
		m.valueInput.View() + "\n\n" +
		helpStyle.Render("🔗 Enter the issue number or description")
}
