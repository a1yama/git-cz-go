package interfaces

type InputHandler interface {
	PromptInput(prompt string) string
}

type CommitExecutor interface {
	Execute(message string)
}
