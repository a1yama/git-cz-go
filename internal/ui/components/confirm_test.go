package components

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewConfirmModel(t *testing.T) {
	// Create model
	model := NewConfirmModel()

	// Verify the model
	view := model.View()
	if view == "" {
		t.Error("View() returned an empty string")
	}

	// By default, confirmation should be true
	if !model.confirmed {
		t.Error("NewConfirmModel() has confirmed=false, want true")
	}

	// Check initialization
	cmd := model.Init()
	if cmd != nil {
		t.Error("Init() returned a non-nil command")
	}
}

func TestConfirmModelUpdate(t *testing.T) {
	// Create model
	model := NewConfirmModel()

	// Test key messages for Y and N
	testCases := []struct {
		name     string
		key      string
		expected bool
		wantCmd  bool
	}{
		{
			name:     "Y key",
			key:      "y",
			expected: true,
			wantCmd:  true,
		},
		{
			name:     "N key",
			key:      "n",
			expected: false,
			wantCmd:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset model for each test
			model = NewConfirmModel()

			keyMsg := tea.KeyMsg{
				Type:  tea.KeyRunes,
				Runes: []rune(tc.key),
			}

			updatedModel, cmd := model.Update(keyMsg)

			// Check that the model was updated
			if updatedModel == nil {
				t.Error("Update() returned nil model")
			}

			// Check updated model type
			_, ok := updatedModel.(ConfirmModel)
			if !ok {
				t.Error("Update() returned a model that is not a ConfirmModel")
			}

			// If the test expects a command, check if one was returned
			if tc.wantCmd && cmd == nil {
				t.Error("Update() should have returned a command but didn't")
			} else if !tc.wantCmd && cmd != nil {
				t.Error("Update() shouldn't have returned a command but did")
			}

			// For yes/no keys, check for expected confirmation state via ConfirmMsg
			if tc.wantCmd {
				msg := executeCmd(t, cmd)
				confirmMsg, ok := msg.(ConfirmMsg)
				if !ok {
					t.Errorf("Command returned %T, want ConfirmMsg", msg)
				}

				if confirmMsg.Confirmed != tc.expected {
					t.Errorf("ConfirmMsg.Confirmed = %v, want %v", confirmMsg.Confirmed, tc.expected)
				}
			}
		})
	}

	// 別途スペースキーのトグルテスト
	t.Run("Space key toggles", func(t *testing.T) {
		model := NewConfirmModel()
		initialState := model.confirmed

		// スペースキーメッセージの作成（文字列として " " を使用）
		spaceMsg := tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune(" "),
		}

		updatedModel, _ := model.Update(spaceMsg)
		newModel, ok := updatedModel.(ConfirmModel)
		if !ok {
			t.Fatal("Update() returned a model that is not a ConfirmModel")
		}

		// トグルされたことを確認
		if newModel.confirmed == initialState {
			t.Errorf("Space key didn't toggle the confirmed state: initial=%v, after=%v",
				initialState, newModel.confirmed)
		}
	})
}

// Helper function to execute a tea.Cmd and return the resulting Msg
func executeCmd(t *testing.T, cmd tea.Cmd) tea.Msg {
	t.Helper()
	if cmd == nil {
		t.Fatal("executeCmd called with nil command")
	}
	return cmd()
}
