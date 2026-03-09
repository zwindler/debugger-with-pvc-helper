package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	primaryColor   = lipgloss.Color("#7C3AED") // Purple
	secondaryColor = lipgloss.Color("#06B6D4") // Cyan
	successColor   = lipgloss.Color("#10B981") // Green
	warningColor   = lipgloss.Color("#F59E0B") // Amber
	errorColor     = lipgloss.Color("#EF4444") // Red
	dimColor       = lipgloss.Color("#6B7280") // Gray
	textColor      = lipgloss.Color("#F9FAFB") // Light

	// Styles
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)

	cursorStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true)

	dimStyle = lipgloss.NewStyle().
			Foreground(dimColor)

	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)

	warningStyle = lipgloss.NewStyle().
			Foreground(warningColor)

	helpStyle = lipgloss.NewStyle().
			Foreground(dimColor).
			MarginTop(1)

	filterPromptStyle = lipgloss.NewStyle().
				Foreground(secondaryColor).
				Bold(true)

	filterTextStyle = lipgloss.NewStyle().
			Foreground(textColor)

	checkboxChecked   = selectedStyle.Render("[x]")
	checkboxUnchecked = dimStyle.Render("[ ]")
)
