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
	prefix := lipgloss.NewStyle().
		Foreground(lipgloss.Color("243")).
		Render(fmt.Sprintf("%d", i.index))
	
	if i.useEmoji && i.emoji != "" {
		return fmt.Sprintf("%s %s %s", prefix, i.emoji, i.type_)
	}
	return fmt.Sprintf("%s %s", prefix, i.type_)
}

// Description returns the description for the list item
func (i commitTypeItem) Description() string { 
	return fmt.Sprintf("• %s", i.description)
}

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
			index:       i + 1, // 1-based indexing for user-friendly display
		}
	}

	// コンパクトなサイズ
	width := 70
	height := len(types) + 2 // より適切な高さ

	// Set up list with custom delegate
	delegate := list.NewDefaultDelegate()
	
	// 項目の高さを1に設定してコンパクトに
	delegate.SetHeight(1)
	delegate.SetSpacing(0)

	// モダンなカラーパレット
	selectedBg := lipgloss.Color("62")    // 青緑
	selectedFg := lipgloss.Color("230")   // 明るい黄色
	normalFg := lipgloss.Color("39")      // 青
	descFg := lipgloss.Color("245")       // グレー

	// Customize delegate styles - 選択時のスタイル
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		Foreground(selectedFg).
		Background(selectedBg).
		Bold(true).
		PaddingLeft(1).
		PaddingRight(1)

	delegate.Styles.SelectedDesc = lipgloss.NewStyle().
		Foreground(selectedFg).
		Background(selectedBg).
		Italic(true).
		PaddingLeft(1).
		PaddingRight(1)

	// 通常時のスタイル
	delegate.Styles.NormalTitle = lipgloss.NewStyle().
		Foreground(normalFg).
		Bold(true).
		PaddingLeft(1)

	delegate.Styles.NormalDesc = lipgloss.NewStyle().
		Foreground(descFg).
		PaddingLeft(1)

	listModel := list.New(items, delegate, width, height)
	listModel.SetShowHelp(false) // ヘルプを非表示にしてコンパクトに
	listModel.SetFilteringEnabled(false) // フィルタリングも無効にして簡潔に
	listModel.SetShowStatusBar(false) // ステータスバーも非表示
	listModel.SetShowTitle(false) // タイトルも非表示

	return CommitTypeModel{
		list: listModel,
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
		// 最適な幅と高さを設定
		width := msg.Width - 4
		if width > 70 {
			width = 70
		}
		height := len(m.list.Items()) + 2
		if height > msg.Height-8 {
			height = msg.Height - 8
		}
		m.list.SetWidth(width)
		m.list.SetHeight(height)
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
	// コンパクトなヒント
	hintStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("243")).
		MarginTop(1).
		Italic(true)
	
	hint := hintStyle.Render("💡 Press number keys (1-9) for quick selection")
	return m.list.View() + "\n" + hint
}
