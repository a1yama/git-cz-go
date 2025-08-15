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
	if cmd := model.Init(); cmd != nil {
		t.Error("Init() should return nil for the new simple implementation")
	}
}

func TestCommitTypeItemTitle(t *testing.T) {
	// This test is no longer relevant since we removed commitTypeItem
	// Instead, test the model's functionality directly

	types := []config.CommitType{
		{Type: "feat", Description: "A new feature", Emoji: "✨"},
	}

	model := NewCommitTypeModel(types, false)
	view := model.View()

	// Check that the view contains expected elements
	if !containsText(view, "feat") {
		t.Errorf("View should contain 'feat', got %q", view)
	}

	if !containsText(view, "A new feature") {
		t.Errorf("View should contain 'A new feature', got %q", view)
	}

	// Test with emoji enabled
	modelWithEmoji := NewCommitTypeModel(types, true)
	viewWithEmoji := modelWithEmoji.View()

	if !containsText(viewWithEmoji, "✨") {
		t.Errorf("View with emoji should contain '✨', got %q", viewWithEmoji)
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

	// Test navigation with arrow keys
	keyMsg := tea.KeyMsg{Type: tea.KeyDown}
	updatedModel, cmd := model.Update(keyMsg)

	// Check that the model was updated
	if updatedModel == nil {
		t.Error("Update() returned nil model")
	}

	// Check that cursor moved
	typedModel, ok := updatedModel.(CommitTypeModel)
	if !ok {
		t.Error("Update() returned a model that is not a CommitTypeModel")
	}

	if typedModel.cursor != 1 {
		t.Errorf("Expected cursor to be 1, got %d", typedModel.cursor)
	}

	// Test numerical key press
	keyMsg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'1'}}
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
