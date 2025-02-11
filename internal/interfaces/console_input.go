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

func (ci *ConsoleInput) PromptInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
