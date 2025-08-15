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
	// 新しい形式では番号は styled prefix で、"1 feat"形式になる
	if title == "" {
		t.Error("Title() without emoji returned empty string")
	}
	// スタイルを含むので、正確な文字列マッチングではなく、含まれる内容をチェック
	if !containsText(title, "1") || !containsText(title, "feat") {
		t.Errorf("Title() without emoji should contain '1' and 'feat', got %q", title)
	}

	// Test with emoji
	item.useEmoji = true
	title = item.Title()
	if !containsText(title, "1") || !containsText(title, "✨") || !containsText(title, "feat") {
		t.Errorf("Title() with emoji should contain '1', '✨' and 'feat', got %q", title)
	}
}

// Helper function to check if styled text contains expected content
func containsText(styled, expected string) bool {
	// Remove ANSI escape sequences and check content
	cleaned := stripANSI(styled)
	return len(cleaned) > 0 && contains(cleaned, expected)
}

// Simple contains check
func contains(s, substr string) bool {
	return len(s) >= len(substr) && findSubstring(s, substr)
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Simple ANSI escape sequence stripper
func stripANSI(s string) string {
	result := ""
	inEscape := false
	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if r == 'm' {
				inEscape = false
			}
			continue
		}
		result += string(r)
	}
	return result
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
