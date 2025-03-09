package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// IsGitRepository checks if the current directory is a git repository
func IsGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}

// GetGitRootDir returns the git repository root directory
func GetGitRootDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// GetBranches returns a list of git branches
func GetBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--format", "%(refname:short)")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(strings.TrimSpace(string(output)), "\n")
	return branches, nil
}

// GetCurrentBranch returns the current git branch
func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// GetStagedFiles returns a list of staged files
func GetStagedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	if files[0] == "" && len(files) == 1 {
		return []string{}, nil
	}
	return files, nil
}

// Commit commits changes with the given message
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// DetectScopes tries to detect scopes from the repository structure
func DetectScopes() ([]string, error) {
	rootDir, err := GetGitRootDir()
	if err != nil {
		return nil, err
	}

	var scopes []string

	// Add common directories as scopes
	for _, dir := range []string{"cmd", "pkg", "internal", "api", "ui", "docs"} {
		path := fmt.Sprintf("%s/%s", rootDir, dir)
		if _, err := os.Stat(path); err == nil {
			// Check if the directory has files tracked by git
			cmd := exec.Command("git", "ls-files", path)
			output, err := cmd.Output()
			if err == nil && len(output) > 0 {
				scopes = append(scopes, dir)
			}
		}
	}

	// Get all git-tracked directories in the root
	cmd := exec.Command("git", "ls-files", "--directory", "--full-name", rootDir)
	output, err := cmd.Output()
	if err != nil {
		return scopes, nil // Return what we have so far
	}

	files := strings.Split(strings.TrimSpace(string(output)), "\n")
	scopeMap := make(map[string]bool)

	for _, file := range files {
		parts := strings.Split(file, "/")
		if len(parts) > 0 {
			scopeMap[parts[0]] = true
		}
	}

	for scope := range scopeMap {
		if !contains(scopes, scope) {
			scopes = append(scopes, scope)
		}
	}

	return scopes, nil
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
