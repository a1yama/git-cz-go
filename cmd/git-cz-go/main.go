package main

import (
	"fmt"
	"os"

	"github.com/a1yama/git-cz-go/internal/config"
	"github.com/a1yama/git-cz-go/internal/git"
	"github.com/a1yama/git-cz-go/internal/ui"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Check if we're in a git repository
	if !git.IsGitRepository() {
		fmt.Fprintln(os.Stderr, "Error: not a git repository (or any of the parent directories)")
		os.Exit(1)
	}

	// Start the TUI
	if err := ui.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
