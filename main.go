package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Commit types based on conventional commits
var commitTypes = []string{
	"feat: A new feature",
	"fix: A bug fix",
	"docs: Documentation changes",
	"style: Code style changes (white-space, formatting, etc.)",
	"refactor: Refactoring code without changing functionality",
	"perf: Performance improvements",
	"test: Adding or updating tests",
	"chore: Miscellaneous tasks",
}

func main() {
	// Select commit type
	fmt.Println("Select a commit type:")
	for i, commitType := range commitTypes {
		fmt.Printf("[%d] %s\n", i+1, commitType)
	}

	var commitTypeIndex int
	fmt.Print("Enter the number of the commit type: ")
	_, err := fmt.Scanln(&commitTypeIndex)
	if err != nil || commitTypeIndex < 1 || commitTypeIndex > len(commitTypes) {
		fmt.Println("Invalid selection.")
		return
	}

	selectedCommitType := strings.Split(commitTypes[commitTypeIndex-1], ":")[0]

	// Enter scope (optional)
	scope := promptInput("Enter the scope of the change (optional): ")
	if scope != "" {
		scope = fmt.Sprintf("(%s)", scope)
	}

	// Enter commit message
	summary := promptInput("Enter a short description of the change: ")
	if summary == "" {
		fmt.Println("Commit message cannot be empty.")
		return
	}

	// Optionally provide a longer description
	body := promptInput("Enter a longer description (optional): ")

	// Construct the commit message
	commitMessage := fmt.Sprintf("%s%s: %s", selectedCommitType, scope, summary)
	if body != "" {
		commitMessage += fmt.Sprintf("\n\n%s", body)
	}

	// Confirm and execute git commit
	fmt.Println("\nGenerated commit message:\n")
	fmt.Println(commitMessage)
	
	confirm := promptInput("Do you want to proceed with this commit? (y/n): ")
	if strings.ToLower(confirm) != "y" {
		fmt.Println("Commit cancelled.")
		return
	}

	executeGitCommit(commitMessage)
}

func promptInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func executeGitCommit(message string) {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing git commit: %v\n", err)
	} else {
		fmt.Println("Commit successfully created.")
	}
}

