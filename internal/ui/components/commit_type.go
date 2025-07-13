package components

import (
	"fmt"
	"github.com/a1yama/git-cz-go/internal/config"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// CommitTypeSelectedMsg is sent when a commit type is selected
type CommitTypeSelectedMsg struct {
	Type string
}

// commitTypeItem represents a commit type in the list
type commitTypeItem struct {
	type_       string
	description string
	emoji       string
	useEmoji    bool
	index       int
}

// FilterValue implements list.Item
func (i commitTypeItem) FilterValue() string { return i.type_ + " " + i.description }

// Title returns the title for the list item
func (i commitTypeItem) Title() string {
	prefix := fmt.Sprintf("%d. ", i.index)
	if i.useEmoji && i.emoji != "" {
		return prefix + i.emoji + "  " + i.type_
	}
	return prefix + i.type_
}

// Description returns the description for the list item
func (i commitTypeItem) Description() string { return i.description }

// CommitTypeModel handles the commit type selection
type CommitTypeModel struct {
	list        list.Model
	stagedFiles []string
}

// NewCommitTypeModel creates a new commit type model
func NewCommitTypeModel(types []config.CommitType, useEmoji bool, stagedFiles []string) CommitTypeModel {
	items := make([]list.Item, len(types))
	for i, t := range types {
		items[i] = commitTypeItem{
			type_:       t.Type,
			description: t.Description,
			emoji:       t.Emoji,
			useEmoji:    useEmoji,
			index:       i + 1, // 1-based indexing for user-friendly display
		}
	}

	// デフォルトのサイズ
	width := 80
	height := 15

	// Set up list with custom delegate
	delegate := list.NewDefaultDelegate()

	// Customize delegate styles to ensure list items are visible
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#0077CC")).
		Bold(true).
		Padding(0, 1)

	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.
		Foreground(lipgloss.Color("#DDDDDD")).
		Background(lipgloss.Color("#0077CC")).
		Padding(0, 1)

	// Make non-selected items more visible too
	delegate.Styles.NormalTitle = delegate.Styles.NormalTitle.
		Foreground(lipgloss.Color("#0077CC")).
		Bold(true).
		Padding(0, 1)

	delegate.Styles.NormalDesc = delegate.Styles.NormalDesc.
		Foreground(lipgloss.Color("#666666")).
		Padding(0, 1)

	listModel := list.New(items, delegate, width, height)
	listModel.Title = "Commit Types"
	listModel.SetShowHelp(true) // Show help text
	listModel.SetFilteringEnabled(true)

	// Customize list styles
	listModel.Styles.Title = lipgloss.NewStyle().
		MarginLeft(2).
		Bold(true).
		Foreground(lipgloss.Color("#0077CC"))

	listModel.Styles.PaginationStyle = lipgloss.NewStyle().
		Padding(0, 2)

	// Make sure the list itself is visible
	listModel.Styles.NoItems = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		MarginLeft(2).
		MarginTop(1)

	// Ensure help text is visible
	listModel.Styles.HelpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#666666")).
		MarginLeft(2).
		MarginBottom(1)

	return CommitTypeModel{
		list:        listModel,
		stagedFiles: stagedFiles,
	}
}

// Init initializes the model
func (m CommitTypeModel) Init() tea.Cmd {
	// Ensure the list is properly initialized
	return tea.Batch(
		m.list.StartSpinner(),
		m.list.SetItems(m.list.Items()),
	)
}

// Update handles updates for the model
func (m CommitTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width - 4)
		m.list.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		// Check for enter to select an item
		if msg.String() == "enter" {
			i, ok := m.list.SelectedItem().(commitTypeItem)
			if ok {
				return m, func() tea.Msg {
					return CommitTypeSelectedMsg{Type: i.type_}
				}
			}
		} else if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		} else {
			// Check if the key is a number
			key := msg.String()
			if key >= "1" && key <= "9" {
				// Convert key to index (0-based)
				index := int(key[0] - '1')

				// Check if the index is valid
				if index >= 0 && index < len(m.list.Items()) {
					// Select the item
					m.list.Select(index)

					// Get the selected item
					i, ok := m.list.SelectedItem().(commitTypeItem)
					if ok {
						return m, func() tea.Msg {
							return CommitTypeSelectedMsg{Type: i.type_}
						}
					}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the model
func (m CommitTypeModel) View() string {
	// Build staged files display
	filesDisplay := ""
	if len(m.stagedFiles) > 0 {
		filesDisplay = "\n\n📁 Staged files:\n"
		for _, file := range m.stagedFiles {
			filesDisplay += fmt.Sprintf("   • %s\n", file)
		}
	}

	// Add a hint about number selection
	hint := "\nTip: You can also select a commit type by pressing its number (1-" + fmt.Sprintf("%d", len(m.list.Items())) + ")"
	return m.list.View() + filesDisplay + hint
}
