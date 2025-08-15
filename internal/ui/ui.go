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
	"github.com/charmbracelet/lipgloss"
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
	StepBody
	StepBreaking
	StepFooter
	StepConfirm
)

// New creates a new UI model
func New(cfg *config.Config) Model {
	// ステップを初期化
	steps := []tea.Model{
		components.NewCommitTypeModel(cfg.Types, cfg.UseEmoji),
		components.NewScopeModel(),
		components.NewSubjectModel(cfg.MaxSubjectLength),
		components.NewBodyModel(),
		components.NewBreakingModel(),
		components.NewFooterModel(),
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
			components.NewScopeModel(),
			components.NewSubjectModel(m.config.MaxSubjectLength),
			components.NewBodyModel(),
			components.NewBreakingModel(),
			components.NewFooterModel(),
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
		isInputFocused := m.activeStep == int(StepScope) ||
			m.activeStep == int(StepSubject) ||
			m.activeStep == int(StepBody) ||
			m.activeStep == int(StepFooter)

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

	case components.ScopeSubmittedMsg:
		m.commitMessage.Scope = msg.Scope
		m.activeStep++
		return m, m.steps[m.activeStep].Init()

	case components.SubjectSubmittedMsg:
		m.commitMessage.Subject = msg.Subject
		m.activeStep++
		return m, m.steps[m.activeStep].Init()

	case components.BodySubmittedMsg:
		m.commitMessage.Body = msg.Body
		m.activeStep++
		return m, m.steps[m.activeStep].Init()

	case components.BreakingSelectedMsg:
		m.commitMessage.Breaking = msg.IsBreaking
		m.activeStep++
		return m, m.steps[m.activeStep].Init()

	case components.FooterSubmittedMsg:
		m.commitMessage.FooterType = msg.FooterType
		m.commitMessage.FooterValue = msg.FooterValue
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

// getStepNames returns the names of all steps
func (m Model) getStepNames() []string {
	return []string{
		"Type",
		"Scope",
		"Subject",
		"Body",
		"Breaking",
		"Footer",
		"Confirm",
	}
}

// getCurrentCommitPreview returns a preview of the current commit message
func (m Model) getCurrentCommitPreview() string {
	if m.commitMessage.Type == "" {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("243")).
			Render("No selections yet...")
	}

	// Build partial commit message
	preview := ""

	// Add emoji if configured
	if m.config.UseEmoji && m.commitMessage.Emoji != "" {
		preview += m.commitMessage.Emoji + " "
	}

	// Add type
	preview += m.commitMessage.Type

	// Add scope if set
	if m.commitMessage.Scope != "" {
		preview += "(" + m.commitMessage.Scope + ")"
	}

	// Add breaking change marker if set
	if m.commitMessage.Breaking {
		preview += "!"
	}

	// Add subject if set
	if m.commitMessage.Subject != "" {
		preview += ": " + m.commitMessage.Subject
	} else if m.activeStep > int(StepScope) {
		preview += ": ..."
	}

	return preview
}

// renderProgressBar renders a progress bar showing completion status
func (m Model) renderProgressBar() string {
	steps := m.getStepNames()
	var parts []string

	for i, stepName := range steps {
		var style lipgloss.Style
		var marker string

		if i < m.activeStep {
			// Completed step
			style = lipgloss.NewStyle().
				Foreground(lipgloss.Color("46")). // Green
				Bold(true)
			marker = "✓"
		} else if i == m.activeStep {
			// Current step
			style = lipgloss.NewStyle().
				Foreground(lipgloss.Color("33")). // Blue
				Bold(true)
			marker = "●"
		} else {
			// Future step
			style = lipgloss.NewStyle().
				Foreground(lipgloss.Color("243")) // Gray
			marker = "○"
		}

		stepText := fmt.Sprintf("%s %s", marker, stepName)
		parts = append(parts, style.Render(stepText))
	}

	return strings.Join(parts, lipgloss.NewStyle().
		Foreground(lipgloss.Color("243")).
		Render(" → "))
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
	case int(StepScope):
		stepTitle = "Denote the scope of this change (optional)"
	case int(StepSubject):
		stepTitle = "Write a short, imperative tense description of the change"
	case int(StepBody):
		stepTitle = "Provide a longer description of the change (optional)"
	case int(StepBreaking):
		stepTitle = "Are there any breaking changes?"
	case int(StepFooter):
		stepTitle = "List any issues or breaking changes (optional)"
	case int(StepConfirm):
		stepTitle = "Confirm your commit message"
	}

	// Create header with title
	header := styles.HeaderStyle.Render("Git Conventional Commit") + "\n\n"

	// Add current commit preview (if we have selections)
	if m.commitMessage.Type != "" {
		previewStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("86")).
			Bold(true).
			MarginBottom(1)

		currentPreview := m.getCurrentCommitPreview()
		header += previewStyle.Render("Current: ") + currentPreview + "\n\n"
	}

	// Add progress bar
	header += m.renderProgressBar() + "\n\n"

	// Add step title
	header += styles.StepTitleStyle.Render(stepTitle) + "\n" +
		styles.DividerStyle.Render(strings.Repeat("─", m.width))

	if m.activeStep == int(StepConfirm) {
		// For confirmation step, add commit message preview
		preview := m.commitMessage.Format()
		header += "\n" + styles.PreviewStyle.Render("Final Preview:") + "\n\n" +
			styles.PreviewContentStyle.Render(preview)
	}

	// Render step content
	content := ""
	if m.activeStep < len(m.steps) {
		content = m.steps[m.activeStep].View()
	}

	helpText := "↑/↓: Navigate • Enter: Select • Esc: Back • Ctrl+C/Q: Quit"
	switch m.activeStep {
	case int(StepType):
		helpText = "↑/↓: Navigate • 1-9: Quick Select • Enter: Select • Esc: Back • Ctrl+C/Q: Quit"
	case int(StepBody):
		helpText = "Type message • Enter on empty or Enter twice: Continue • Ctrl+D: Continue • Esc: Back • Ctrl+C: Quit"
	case int(StepBreaking):
		helpText = "↑/↓: Navigate • Y/N: Quick Select • Enter: Select • Esc: Back • Ctrl+C: Quit"
	}

	return fmt.Sprintf("%s\n\n%s\n\n%s",
		header,
		content,
		styles.HelpStyle.Render(helpText),
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
