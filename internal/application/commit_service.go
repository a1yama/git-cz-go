package application

import (
	"fmt"
	"os"
	"strings"
	"github.com/a1yama/git-cz-go/internal/domain"
	"github.com/a1yama/git-cz-go/internal/interfaces"
	"github.com/manifoldco/promptui"
)

type CommitService struct {
	inputHandler    interfaces.InputHandler
	commitExecutor  interfaces.CommitExecutor
}

func NewCommitService(inputHandler interfaces.InputHandler, commitExecutor interfaces.CommitExecutor) *CommitService {
	return &CommitService{
		inputHandler:   inputHandler,
		commitExecutor: commitExecutor,
	}
}

func (s *CommitService) Run() {
	commitTypes := []string{
		"feat: A new feature",
		"fix: A bug fix",
		"docs: Documentation changes",
		"style: Code style changes (white-space, formatting, etc.)",
		"refactor: Refactoring code without changing functionality",
		"perf: Performance improvements",
		"test: Adding or updating tests",
		"chore: Miscellaneous tasks",
	}

	selectedType, err := s.selectCommitTypePrompt(commitTypes)
	if err != nil {
		fmt.Println("Prompt cancelled.")
		os.Exit(1) // Exit the program if Ctrl+C is pressed
	}

	scope := s.inputHandler.PromptInput("Enter the scope of the change (optional): ")
	summary := s.inputHandler.PromptInput("Enter a short description of the change: ")
	if summary == "" {
		fmt.Println("Commit message cannot be empty.")
		return
	}
	body := s.inputHandler.PromptInput("Enter a longer description (optional): ")

	commit := domain.NewCommit(selectedType, scope, summary, body)
	message := commit.ConstructMessage()

	fmt.Println("\nGenerated commit message:\n")
	fmt.Println(message)

	confirm := s.inputHandler.PromptInput("Do you want to proceed with this commit? (y/n): ")
	if strings.ToLower(confirm) == "y" {
		s.commitExecutor.Execute(message)
	} else {
		fmt.Println("Commit cancelled.")
	}
}

func (s *CommitService) selectCommitTypePrompt(commitTypes []string) (string, error) {
	prompt := promptui.Select{
		Label:             "Select a commit type",
		Items:             commitTypes,
		Size:              len(commitTypes), // Prevent scrolling
		HideHelp:          true,             // Hide navigation help text
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return strings.Split(result, ":")[0], nil
}
