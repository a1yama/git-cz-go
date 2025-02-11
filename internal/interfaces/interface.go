package interfaces

type InputHandler interface {
	SelectCommitType(commitTypes []string) string
	PromptInput(prompt string) string
}

type CommitExecutor interface {
	Execute(message string)
}