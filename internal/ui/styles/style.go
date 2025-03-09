package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	primaryColor   = lipgloss.Color("69")
	secondaryColor = lipgloss.Color("99")
	accentColor    = lipgloss.Color("208")
	textColor      = lipgloss.Color("252")
	mutedColor     = lipgloss.Color("240")
	successColor   = lipgloss.Color("76")
	errorColor     = lipgloss.Color("160")
	warningColor   = lipgloss.Color("214")
	infoColor      = lipgloss.Color("75")

	// Base styles
	BaseStyle = lipgloss.NewStyle().
			Foreground(textColor)

	// Header styles
	HeaderStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			MarginBottom(1).
			Padding(0, 1)

	ProgressStyle = lipgloss.NewStyle().
			Background(secondaryColor).
			Foreground(lipgloss.Color("255")).
			Padding(0, 1).
			MarginLeft(1)

	StepTitleStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Bold(true)

	DividerStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	// Preview styles
	PreviewStyle = lipgloss.NewStyle().
			Foreground(infoColor).
			Bold(true)

	PreviewContentStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(mutedColor).
				Padding(1, 2).
				MarginTop(1)

	// Help style
	HelpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true)

	// Input styles
	FocusedStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	BlurredStyle = lipgloss.NewStyle().
			Foreground(mutedColor)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(warningColor)

	InfoStyle = lipgloss.NewStyle().
			Foreground(infoColor)
)
