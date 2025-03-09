package components

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// FooterTypeSubmittedMsg is sent when a footer type is selected
type FooterTypeSubmittedMsg struct {
	FooterType string
}

// FooterValueSubmittedMsg is sent when a footer value is submitted
type FooterValueSubmittedMsg struct {
	Value string
}

// footerTypeItem represents a footer type in the list
type footerTypeItem struct {
	name        string
	description string
}

// FilterValue implements list.Item
func (i footerTypeItem) FilterValue() string { return i.name + " " + i.description }

// Title returns the title for the list item
func (i footerTypeItem) Title() string { return i.name }

// Description returns the description for the list item
func (i footerTypeItem) Description() string { return i.description }

// FooterTypeModel handles the footer type selection
type FooterTypeModel struct {
	list list.Model
}

// NewFooterTypeModel creates a new footer type model
func NewFooterTypeModel() FooterTypeModel {
	// Define common footer types
	footerTypes := []list.Item{
		footerTypeItem{name: "", description: "No footer (Skip)"},
		footerTypeItem{name: "BREAKING CHANGE", description: "Introduces a breaking API change"},
		footerTypeItem{name: "Fixes", description: "This change fixes a specific issue"},
		footerTypeItem{name: "Closes", description: "This change closes a specific issue"},
		footerTypeItem{name: "Refs", description: "References an issue or PR"},
		footerTypeItem{name: "See", description: "References external information"},
		footerTypeItem{name: "DEPRECATED", description: "Marks deprecated functionality"},
	}

	// Set up list
	listModel := list.New(footerTypes, list.NewDefaultDelegate(), 0, 0)
	listModel.Title = "Footer Types"
	listModel.SetShowHelp(false)
	listModel.SetFilteringEnabled(true)
	listModel.Styles.Title = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	listModel.Styles.PaginationStyle = lipgloss.NewStyle().Padding(0, 2)

	return FooterTypeModel{
		list: listModel,
	}
}

// Init initializes the model
func (m FooterTypeModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the model
func (m FooterTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width - 4)
		m.list.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		// Check for enter to select an item
		if msg.String() == "enter" {
			i, ok := m.list.SelectedItem().(footerTypeItem)
			if ok {
				return m, func() tea.Msg {
					return FooterTypeSubmittedMsg{FooterType: i.name}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the model
func (m FooterTypeModel) View() string {
	return m.list.View()
}

// FooterValueModel handles the footer value input
type FooterValueModel struct {
	textInput textinput.Model
}

// NewFooterValueModel creates a new footer value model
func NewFooterValueModel() FooterValueModel {
	ti := textinput.New()
	ti.Placeholder = "Enter footer value (e.g. #123, description of breaking change, etc.)"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 50

	return FooterValueModel{
		textInput: ti,
	}
}

// Init initializes the model
func (m FooterValueModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles updates for the model
func (m FooterValueModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			return m, func() tea.Msg {
				return FooterValueSubmittedMsg{Value: m.textInput.Value()}
			}
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View renders the model
func (m FooterValueModel) View() string {
	return m.textInput.View()
}
