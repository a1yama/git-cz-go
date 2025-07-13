package components

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewSubjectModel(t *testing.T) {
	// Create model
	model := NewSubjectModel(100, []string{})

	// Verify the model
	view := model.View()
	if view == "" {
		t.Error("View() returned an empty string")
	}

	// Check initialization
	cmd := model.Init()
	if cmd == nil {
		t.Error("Init() returned a nil command")
	}
}

func TestSubjectModelUpdate(t *testing.T) {
	// Create model with staged files
	stagedFiles := []string{"file1.txt", "src/main.go"}
	model := NewSubjectModel(100, stagedFiles)

	// Test window size message
	windowSizeMsg := tea.WindowSizeMsg{Width: 100, Height: 50}
	updatedModel, _ := model.Update(windowSizeMsg) // cmd変数を使わない（_で捨てる）

	// Check that the model was updated
	if updatedModel == nil {
		t.Error("Update() returned nil model")
	}

	// Check that the model is still a SubjectModel
	_, ok := updatedModel.(SubjectModel)
	if !ok {
		t.Error("Update() returned a model that is not a SubjectModel")
	}

	// Check validation
	if model.validInput {
		t.Error("New model should have validInput=false")
	}
}

func TestSubjectModelViewWithStagedFiles(t *testing.T) {
	// Test with no staged files
	model := NewSubjectModel(100, []string{})
	view := model.View()
	if strings.Contains(view, "Staged files:") {
		t.Error("View() should not display staged files section when there are no staged files")
	}

	// Test with staged files
	stagedFiles := []string{"file1.txt", "src/main.go", "README.md"}
	model = NewSubjectModel(100, stagedFiles)
	view = model.View()

	if !strings.Contains(view, "Staged files:") {
		t.Error("View() should display staged files section when there are staged files")
	}

	for _, file := range stagedFiles {
		if !strings.Contains(view, file) {
			t.Errorf("View() should display staged file %s", file)
		}
	}
}
