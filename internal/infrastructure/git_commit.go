package infrastructure

import (
	"fmt"
	"os"
	"os/exec"
)

type GitCommitExecutor struct {}

func NewGitCommitExecutor() *GitCommitExecutor {
	return &GitCommitExecutor{}
}

func (e *GitCommitExecutor) Execute(message string) {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing git commit: %v\n", err)
	} else {
		fmt.Println("Commit successfully created.")
	}
}