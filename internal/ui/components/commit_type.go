package components

import (
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
}

// FilterValue implements list.Item
func (i commitTypeItem) FilterValue() string { return i.type_ + " " + i.description }

// Title returns the title for the list item
func (i commitTypeItem) Title() string {
	if i.useEmoji && i.emoji != "" {
		return i.emoji + "  " + i.type_
	}
	return i.type_
}

// Description returns the description for the list item
func (i commitTypeItem) Description() string { return i.description }

// CommitTypeModel handles the commit type selection
type CommitTypeModel struct {
	list list.Model
}

// NewCommitTypeModel creates a new commit type model
func NewCommitTypeModel(types []config.CommitType, useEmoji bool) CommitTypeModel {
	items := make([]list.Item, len(types))
	for i, t := range types {
		items[i] = commitTypeItem{
			type_:       t.Type,
			description: t.Description,
			emoji:       t.Emoji,
			useEmoji:    useEmoji,
		}
	}

	// デフォルトのサイズ
	width := 80
	height := 15

	// Set up list
	listModel := list.New(items, list.NewDefaultDelegate(), width, height)
	listModel.Title = "Commit Types"
	listModel.SetShowHelp(false)
	listModel.SetFilteringEnabled(true)
	listModel.Styles.Title = lipgloss.NewStyle().MarginLeft(2).Bold(true)
	listModel.Styles.PaginationStyle = lipgloss.NewStyle().Padding(0, 2)

	return CommitTypeModel{
		list: listModel,
	}
}

// Init initializes the model
func (m CommitTypeModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the model
func (m CommitTypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width - 4)
		m.list.SetHeight(msg.Height - 10)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i, ok := m.list.SelectedItem().(commitTypeItem)
			if ok {
				// このメッセージを親モデルに送信
				return m, func() tea.Msg {
					return CommitTypeSelectedMsg{Type: i.type_}
				}
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the model
func (m CommitTypeModel) View() string {
	return m.list.View()
}
