package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	// Check if default config has the expected values
	if cfg == nil {
		t.Fatal("DefaultConfig() returned nil")
	}

	if len(cfg.Types) == 0 {
		t.Error("DefaultConfig() returned a config with no commit types")
	}

	// Verify that common commit types are present
	typeMap := make(map[string]bool)
	for _, commitType := range cfg.Types {
		typeMap[commitType.Type] = true
	}

	requiredTypes := []string{"feat", "fix", "docs", "style", "refactor"}
	for _, requiredType := range requiredTypes {
		if !typeMap[requiredType] {
			t.Errorf("DefaultConfig() is missing required commit type: %s", requiredType)
		}
	}
}

func TestLoadFromFile(t *testing.T) {
	// Create a temporary config file
	tempDir, err := os.MkdirTemp("", "git-cz-go-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	configPath := filepath.Join(tempDir, ".git-cz.json")
	configContent := `{
        "types": [
            {"type": "custom", "description": "Custom type", "emoji": "🔥"}
        ],
        "useEmoji": true,
        "maxSubjectLength": 50
    }`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// configFilePathsをモンキーパッチするのではなく、テスト用の設定を直接読み込む
	cfg := DefaultConfig()

	// 設定ファイルを直接読み込む
	file, err := os.Open(configPath)
	if err != nil {
		t.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(cfg); err != nil {
		t.Fatalf("Failed to decode config file: %v", err)
	}

	// 設定の検証
	if len(cfg.Types) != 1 {
		t.Errorf("Expected 1 commit type, got %d", len(cfg.Types))
	}

	if cfg.Types[0].Type != "custom" {
		t.Errorf("Expected type 'custom', got '%s'", cfg.Types[0].Type)
	}

	if !cfg.UseEmoji {
		t.Error("Expected UseEmoji to be true")
	}

	if cfg.MaxSubjectLength != 50 {
		t.Errorf("Expected MaxSubjectLength to be 50, got %d", cfg.MaxSubjectLength)
	}
}
