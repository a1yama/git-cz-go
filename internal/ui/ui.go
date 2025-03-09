package ui

import (
	"fmt"
	"strings"

	"github.com/a1yama/git-cz-go/internal/config"
	"github.com/a1yama/git-cz-go/internal/git"
	"github.com/a1yama/git-cz-go/internal/model"
	"github.com/a1yama/git-cz-go/internal/ui/components"
	"github.com/a1yama/git-cz-go/internal/ui/styles"
	"github.com/charmbracelet/bubbles/key"
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
	StepSubject
	StepConfirm
)

// New creates a new UI model
func New(cfg *config.Config) Model {
	// ステップを初期化
	steps := []tea.Model{
		components.NewCommitTypeModel(cfg.Types, cfg.UseEmoji),
		components.NewSubjectModel(cfg.MaxSubjectLength),
		components.NewConfirmModel(),
	}

	return Model{
		config:     cfg,
		activeStep: 0,
		steps:      steps, // 初期化したステップを設定
		ready:      false,
	}
}

// Init関数も修正
func (m Model) Init() tea.Cmd {
	// ステップがすでに初期化されていることを確認
	if len(m.steps) == 0 {
		// 万が一ステップが空の場合は、ここで初期化
		m.steps = []tea.Model{
			components.NewCommitTypeModel(m.config.Types, m.config.UseEmoji),
			components.NewSubjectModel(m.config.MaxSubjectLength),
			components.NewConfirmModel(),
		}
	}

	// 最初のステップの初期化コマンドを返す
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
		// テキスト入力フォーカス中はグローバルショートカットを無効化
		isInputFocused := m.activeStep == int(StepSubject)

		// Global keybindings（テキスト入力中は無効）
		if !isInputFocused {
			switch {
			case key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c", "q"))):
				return m, tea.Quit

			case key.Matches(msg, key.NewBinding(key.WithKeys("esc"))):
				if m.activeStep > 0 {
					m.activeStep--
					return m, m.steps[m.activeStep].Init()
				}
				return m, tea.Quit
			}
		} else {
			// テキスト入力中はCtrl+Cのみ終了として扱う
			if key.Matches(msg, key.NewBinding(key.WithKeys("ctrl+c"))) {
				return m, tea.Quit
			}
			// Escキーはテキスト入力中でも前のステップに戻る
			if key.Matches(msg, key.NewBinding(key.WithKeys("esc"))) {
				if m.activeStep > 0 {
					m.activeStep--
					return m, m.steps[m.activeStep].Init()
				}
				return m, tea.Quit
			}
		}

	// 他のメッセージハンドリング...
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
		m.activeStep++
		if m.activeStep >= len(m.steps) {
			return m, tea.Quit
		}
		return m, m.steps[m.activeStep].Init()

	case components.SubjectSubmittedMsg:
		m.commitMessage.Subject = msg.Subject
		m.activeStep++
		return m, m.steps[m.activeStep].Init()

	case components.ConfirmMsg:
		if msg.Confirmed {
			commitMsg := m.commitMessage.Format()
			return m, tea.Sequence(
				commitCmd(commitMsg),
				tea.Quit,
			)
		}
		return m, tea.Quit
	}

	// Pass the message to the current step
	if m.activeStep < len(m.steps) {
		updatedStep, cmd := m.steps[m.activeStep].Update(msg)
		m.steps[m.activeStep] = updatedStep
		cmds = append(cmds, cmd)
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

	var stepTitle string
	switch m.activeStep {
	case int(StepType):
		stepTitle = "Select the type of change that you're committing"
	case int(StepSubject):
		stepTitle = "Write a short, imperative tense description of the change"
	case int(StepConfirm):
		stepTitle = "Confirm your commit message"
	}

	// Display progress
	progress := fmt.Sprintf(" %d/%d ", m.activeStep+1, len(m.steps))

	header := styles.HeaderStyle.Render("Git Conventional Commit") +
		styles.ProgressStyle.Render(progress) +
		"\n\n" +
		styles.StepTitleStyle.Render(stepTitle) +
		"\n" +
		styles.DividerStyle.Render(strings.Repeat("─", m.width))

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

	return fmt.Sprintf("%s\n\n%s\n\n%s",
		header,
		content,
		styles.HelpStyle.Render("↑/↓: Navigate • Enter: Select • Esc: Back • Ctrl+C/Q: Quit"),
	)
}

// Run runs the UI
func Run(cfg *config.Config) error {
	p := tea.NewProgram(New(cfg), tea.WithAltScreen())
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
