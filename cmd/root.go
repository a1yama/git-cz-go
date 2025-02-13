package cmd

import (
	"github.com/a1yama/git-cz-go/internal/application"
	"github.com/a1yama/git-cz-go/internal/infrastructure"
	"github.com/a1yama/git-cz-go/internal/interfaces"
)

// Execute はプログラムのエントリーポイント
func Execute() {
	commitService := application.NewCommitService(
		interfaces.NewConsoleInput(),
		infrastructure.NewGitCommitExecutor(),
	)
	commitService.Run()
}