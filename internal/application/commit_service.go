package application

import (
	"fmt"
	"strings"
	"github.com/a1yama/git-cz-go/internal/domain"
	"github.com/a1yama/git-cz-go/internal/interfaces"  // インターフェースのインポート
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

	selectedType := s.inputHandler.SelectCommitType(commitTypes)
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
