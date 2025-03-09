package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// setupGitRepo creates a temporary Git repository for testing
func setupGitRepo(t *testing.T) string {
	t.Helper()

	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "git-cz-go-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}

	// Initialize Git repository
	cmd := exec.Command("git", "init")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		os.RemoveAll(tempDir)
		t.Fatalf("Failed to initialize git repository: %v", err)
	}

	// Configure Git user for commits
	cmd = exec.Command("git", "config", "user.name", "Test User")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		os.RemoveAll(tempDir)
		t.Fatalf("Failed to configure git user name: %v", err)
	}

	cmd = exec.Command("git", "config", "user.email", "test@example.com")
	cmd.Dir = tempDir
	if err := cmd.Run(); err != nil {
		os.RemoveAll(tempDir)
		t.Fatalf("Failed to configure git user email: %v", err)
	}

	return tempDir
}

func TestIsGitRepository(t *testing.T) {
	// Create a temporary Git repository
	tempDir := setupGitRepo(t)
	defer os.RemoveAll(tempDir)

	// Change to the repository directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to repository directory: %v", err)
	}

	// Test IsGitRepository
	if !IsGitRepository() {
		t.Error("IsGitRepository() = false, want true")
	}

	// Test in a non-Git directory
	nonGitDir, err := os.MkdirTemp("", "non-git")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(nonGitDir)

	if err := os.Chdir(nonGitDir); err != nil {
		t.Fatalf("Failed to change to non-git directory: %v", err)
	}

	if IsGitRepository() {
		t.Error("IsGitRepository() = true, want false")
	}
}

func TestGetGitRootDir(t *testing.T) {
	// Create a temporary Git repository
	tempDir := setupGitRepo(t)
	defer os.RemoveAll(tempDir)

	// Change to the repository directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to repository directory: %v", err)
	}

	// Create a subdirectory
	subDir := filepath.Join(tempDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Change to the subdirectory
	if err := os.Chdir(subDir); err != nil {
		t.Fatalf("Failed to change to subdirectory: %v", err)
	}

	// Test GetGitRootDir
	rootDir, err := GetGitRootDir()
	if err != nil {
		t.Fatalf("GetGitRootDir() failed: %v", err)
	}

	// Normalize paths for comparison
	expectedPath, err := filepath.EvalSymlinks(tempDir)
	if err != nil {
		t.Fatalf("Failed to evaluate symlinks for temp directory: %v", err)
	}

	actualPath, err := filepath.EvalSymlinks(rootDir)
	if err != nil {
		t.Fatalf("Failed to evaluate symlinks for root directory: %v", err)
	}

	if actualPath != expectedPath {
		t.Errorf("GetGitRootDir() = %q, want %q", actualPath, expectedPath)
	}
}

func TestCommit(t *testing.T) {
	// Skip in CI environments
	if os.Getenv("CI") != "" {
		t.Skip("Skipping in CI environment")
	}

	// Create a temporary Git repository
	tempDir := setupGitRepo(t)
	defer os.RemoveAll(tempDir)

	// Change to the repository directory
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to repository directory: %v", err)
	}

	// Create a test file
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Stage the file
	cmd := exec.Command("git", "add", "test.txt")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to stage file: %v", err)
	}

	// Commit the file
	message := "test: add test file"
	if err := Commit(message); err != nil {
		t.Fatalf("Commit() failed: %v", err)
	}

	// Verify the commit
	cmd = exec.Command("git", "log", "-1", "--pretty=%B")
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to get commit message: %v", err)
	}

	// Trim the output to handle different git log formats
	commitMessage := string(output)
	if commitMessage == "" || commitMessage != message+"\n" && commitMessage != message+"\n\n" {
		t.Errorf("Commit message = %q, want %q", commitMessage, message)
	}
}
