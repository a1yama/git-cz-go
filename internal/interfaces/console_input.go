package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleInput struct {}

func NewConsoleInput() *ConsoleInput {
	return &ConsoleInput{}
}

func (ci *ConsoleInput) SelectCommitType(commitTypes []string) string {
	fmt.Println("Select a commit type:")
	for i, commitType := range commitTypes {
		fmt.Printf("[%d] %s\n", i+1, commitType)
	}

	var commitTypeIndex int
	fmt.Print("Enter the number of the commit type: ")
	_, err := fmt.Scanln(&commitTypeIndex)
	if err != nil || commitTypeIndex < 1 || commitTypeIndex > len(commitTypes) {
		fmt.Println("Invalid selection.")
		os.Exit(1)
	}
	return strings.Split(commitTypes[commitTypeIndex-1], ":")[0]
}

func (ci *ConsoleInput) PromptInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
