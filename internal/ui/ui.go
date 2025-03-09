package ui

import (
	"fmt"
	"strings"

	"github.com/a1yama/git-cz-go/internal/config"
	"github.com/a1yama/git-cz-go/internal/git"
	"github.com/a1yama/git-cz-go/internal/model"
	"github.com/a1yama/git-cz-go/internal/ui/components"
	"github.com/a1yama/git-cz-go/internal/ui/styles"
	tea "github.com/charmbracelet/bubbletea"
)

// Model is the main UI model
type Model struct {
	config        *config.Config
	commitMessage model.CommitMessage
	activeStep    int
	steps         []tea.Model
	width         int
	height        int
	ready         bool
	err           error
}

// Step represents a commit message input step
type Step int

const (
	StepType Step = iota
	StepScope
	StepSubject
	StepBreaking
	StepBody
	StepFooterType
	StepFooterValue
	StepConfirm
)

// New creates a new UI model
func New(cfg *config.Config) Model {
	// ここでステップを初期化
	var steps []tea.Model
	steps = append(steps, components.NewCommitTypeModel(cfg.Types, cfg.UseEmoji))

	if !cfg.SkipScope {
		scopeModel := components.NewScopeModel()
		scopes, _ := git.DetectScopes()
		scopeModel.SetSuggestions(scopes)
		steps = append(steps, scopeModel)
	}

	// 他のステップも追加
	steps = append(steps, components.NewSubjectModel(cfg.MaxSubjectLength))
	steps = append(steps, components.NewBreakingModel())

	if !cfg.SkipBody {
		steps = append(steps, components.NewBodyModel(cfg.MaxBodyLineLength))
	}

	if !cfg.SkipFooter {
		steps = append(steps, components.NewFooterTypeModel())
		steps = append(steps, components.NewFooterValueModel())
	}

	// 確認ステップ
	steps = append(steps, components.NewConfirmModel())

	return Model{
		config:     cfg,
		activeStep: 0,
		steps:      steps, // ここで初期化済みのステップを設定
		ready:      false,
	}
}

// Init initializes the UI
func (m Model) Init() tea.Cmd {
	if len(m.steps) == 0 {
		return nil
	}
	return m.steps[0].Init()
}

// Update handles UI updates
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

	case tea.KeyMsg:
		// Global keybindings
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			if m.activeStep > 0 {
				m.activeStep--
				return m, m.steps[m.activeStep].Init()
			}
			return m, tea.Quit
		}

	// 以下のメッセージハンドリングを確認
	case components.CommitTypeSelectedMsg:
		m.commitMessage.Type = msg.Type
		if m.config.UseEmoji {
			for _, t := range m.config.Types {
				if t.Type == msg.Type {
					m.commitMessage.Emoji = t.Emoji
					break
				}
			}
		}
		// 次のステップに進む
		m.activeStep++
		if m.activeStep < len(m.steps) {
			return m, m.steps[m.activeStep].Init()
		}

	case components.ScopeSubmittedMsg:
		m.commitMessage.Scope = msg.Scope
		m.activeStep++
		if m.activeStep < len(m.steps) {
			return m, m.steps[m.activeStep].Init()
		}

		// 他のメッセージハンドリングも同様...
	}

	// 現在のステップにメッセージを渡す
	if m.activeStep < len(m.steps) {
		newModel, cmd := m.steps[m.activeStep].Update(msg)
		m.steps[m.activeStep] = newModel
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

// View renders the UI
func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	if m.err != nil {
		return fmt.Sprintf("Error: %v", m.err)
	}

	// Display progress
	progress := fmt.Sprintf(" %d/%d ", m.activeStep+1, len(m.steps))

	header := styles.HeaderStyle.Render("Git Conventional Commit") +
		styles.ProgressStyle.Render(progress) +
		"\n\n" +
		styles.StepTitleStyle.Render("Select the type of change that you're committing") +
		"\n" +
		styles.DividerStyle.Render(strings.Repeat("─", max(m.width, 80)))

	if m.activeStep == int(StepConfirm) {
		// For confirmation step, add commit message preview
		preview := m.commitMessage.Format()
		header += "\n" + styles.PreviewStyle.Render("Preview:") + "\n\n" +
			styles.PreviewContentStyle.Render(preview)
	}

	// Render step content
	content := ""
	if m.activeStep < len(m.steps) {
		content = m.steps[m.activeStep].View()
	}

	// フォールバック - リストが表示されない場合
	if content == "" && m.activeStep == int(StepType) {
		content = "Available types:\n"
		for _, t := range m.config.Types {
			content += fmt.Sprintf("  %s - %s\n", t.Type, t.Description)
		}
	}

	return fmt.Sprintf("%s\n\n%s\n\n%s",
		header,
		content,
		styles.HelpStyle.Render("↑/↓: Navigate • Enter: Select • Esc: Back • Ctrl+C/Q: Quit"),
	)
}

// Helper function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Run runs the UI
func Run(cfg *config.Config) error {
	initialModel := New(cfg)
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// commitCmd creates a command for git commit
func commitCmd(message string) tea.Cmd {
	return func() tea.Msg {
		if err := git.Commit(message); err != nil {
			return err
		}
		return nil
	}
}
