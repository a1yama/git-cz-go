package components

import (
	"testing"

	"github.com/a1yama/git-cz-go/internal/config"
	tea "github.com/charmbracelet/bubbletea"
)

func TestNewCommitTypeModel(t *testing.T) {
	// Create test commit types
	types := []config.CommitType{
		{Type: "feat", Description: "A new feature", Emoji: "✨"},
		{Type: "fix", Description: "A bug fix", Emoji: "🐛"},
	}

	// Create model with emoji disabled
	model := NewCommitTypeModel(types, false)

	// Verify the model
	view := model.View()
	if view == "" {
		t.Error("View() returned an empty string")
	}

	// Check if the model initializes correctly
	if cmd := model.Init(); cmd == nil {
		t.Error("Init() returned a nil command")
	}
}

func TestCommitTypeItemTitle(t *testing.T) {
	// Test without emoji
	item := commitTypeItem{
		type_:       "feat",
		description: "A new feature",
		emoji:       "✨",
		useEmoji:    false,
		index:       1,
	}

	title := item.Title()
	expected := "1. feat"
	if title != expected {
		t.Errorf("Title() without emoji returned %q, expected %q", title, expected)
	}

	// Test with emoji
	item.useEmoji = true
	title = item.Title()
	expected = "1. ✨  feat"
	if title != expected {
		t.Errorf("Title() with emoji returned %q, expected %q", title, expected)
	}
}

func TestCommitTypeModelUpdate(t *testing.T) {
	// Create test commit types
	types := []config.CommitType{
		{Type: "feat", Description: "A new feature", Emoji: "✨"},
		{Type: "fix", Description: "A bug fix", Emoji: "🐛"},
	}

	// Create model
	model := NewCommitTypeModel(types, false)

	// Test window size message
	windowSizeMsg := tea.WindowSizeMsg{Width: 100, Height: 50}
	updatedModel, cmd := model.Update(windowSizeMsg)

	// Check that the model was updated
	if updatedModel == nil {
		t.Error("Update() returned nil model")
	}

	// Check that no command was returned
	if cmd != nil {
		t.Error("Update() with WindowSizeMsg returned a non-nil command")
	}

	// Check that the model is still a CommitTypeModel
	_, ok := updatedModel.(CommitTypeModel)
	if !ok {
		t.Error("Update() returned a model that is not a CommitTypeModel")
	}

	// Test numerical key press
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}}
	updatedModel, cmd = model.Update(keyMsg)

	// Check that the model was updated
	if updatedModel == nil {
		t.Error("Update() returned nil model")
	}

	// Check that a command was returned (for selecting the item)
	if cmd == nil {
		t.Error("Update() with numerical key press returned a nil command")
	}
}
