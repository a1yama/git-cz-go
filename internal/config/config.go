package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// CommitType represents a conventional commit type with description
type CommitType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Emoji       string `json:"emoji,omitempty"`
}

// Config holds the configuration for git-cz-go
type Config struct {
	Types             []CommitType `json:"types"`
	UseEmoji          bool         `json:"useEmoji"`
	CustomTemplate    string       `json:"customTemplate,omitempty"`
	SkipScope         bool         `json:"skipScope"`
	SkipBody          bool         `json:"skipBody"`
	SkipBreakingBody  bool         `json:"skipBreakingBody"`
	SkipFooter        bool         `json:"skipFooter"`
	MaxSubjectLength  int          `json:"maxSubjectLength"`
	MaxBodyLineLength int          `json:"maxBodyLineLength"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Types: []CommitType{
			{Type: "feat", Description: "A new feature", Emoji: "✨"},
			{Type: "fix", Description: "A bug fix", Emoji: "🐛"},
			{Type: "docs", Description: "Documentation only changes", Emoji: "📚"},
			{Type: "style", Description: "Changes that do not affect the meaning of the code", Emoji: "💎"},
			{Type: "refactor", Description: "A code change that neither fixes a bug nor adds a feature", Emoji: "📦"},
			{Type: "perf", Description: "A code change that improves performance", Emoji: "🚀"},
			{Type: "test", Description: "Adding missing tests or correcting existing tests", Emoji: "🚨"},
			{Type: "build", Description: "Changes that affect the build system or external dependencies", Emoji: "🛠"},
			{Type: "ci", Description: "Changes to our CI configuration files and scripts", Emoji: "⚙️"},
			{Type: "chore", Description: "Other changes that don't modify src or test files", Emoji: "♻️"},
			{Type: "revert", Description: "Reverts a previous commit", Emoji: "🗑"},
		},
		UseEmoji:          false,
		SkipScope:         false,
		SkipBody:          false,
		SkipBreakingBody:  false,
		SkipFooter:        false,
		MaxSubjectLength:  100,
		MaxBodyLineLength: 100,
	}
}

// configFilePaths returns a list of possible config file locations
func configFilePaths() ([]string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	// Current directory and user's home directory
	return []string{
		"./.git-cz.json",
		filepath.Join(home, ".git-cz.json"),
		filepath.Join(home, ".config", "git-cz", "config.json"),
	}, nil
}

// Load loads the configuration from disk
func Load() (*Config, error) {
	config := DefaultConfig()

	paths, err := configFilePaths()
	if err != nil {
		return config, err
	}

	// Try to load from each path
	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			file, err := os.Open(path)
			if err != nil {
				continue
			}
			defer file.Close()

			if err := json.NewDecoder(file).Decode(config); err == nil {
				return config, nil
			}
		}
	}

	// Return default config if no config file found
	return config, nil
}

// Save saves the configuration to disk
func (c *Config) Save() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(home, ".config", "git-cz")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	configPath := filepath.Join(configDir, "config.json")
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}
