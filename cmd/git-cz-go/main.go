package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/a1yama/git-cz-go/internal/config"
	"github.com/a1yama/git-cz-go/internal/git"
	"github.com/a1yama/git-cz-go/internal/ui"
)

func main() {
	// Define command-line flags
	showTypes := flag.Bool("types", false, "Show available commit types")
	flag.Parse()

	// Load config
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// If --types flag is specified, show commit types and exit
	if *showTypes {
		fmt.Println("Available commit types:")
		fmt.Println("----------------------")
		for _, t := range cfg.Types {
			emoji := ""
			if cfg.UseEmoji && t.Emoji != "" {
				emoji = t.Emoji + "  "
			}
			fmt.Printf("%s%s: %s\n", emoji, t.Type, t.Description)
		}
		return
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
