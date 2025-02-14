package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	tabStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(0, 1).
			Margin(0, 1)

	activeTabStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#F25D94")).
			Background(lipgloss.Color("#F25D94")).
			Foreground(lipgloss.Color("#FFF7DB")).
			Padding(0, 1).
			Margin(0, 1)

	inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(0, 1).
			Width(60)

	buttonStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#7D56F4")).
			Foreground(lipgloss.Color("#FAFAFA")).
			Padding(0, 2).
			Margin(0, 1)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575"))

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F87"))

	methodStyle = map[string]lipgloss.Style{
		"GET":    lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).Bold(true),
		"POST":   lipgloss.NewStyle().Foreground(lipgloss.Color("#F25D94")).Bold(true),
		"PUT":    lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Bold(true),
		"DELETE": lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F87")).Bold(true),
		"PATCH":  lipgloss.NewStyle().Foreground(lipgloss.Color("#FFBD2E")).Bold(true),
		"HEAD":   lipgloss.NewStyle().Foreground(lipgloss.Color("#8B8B8B")).Bold(true),
	}

	panelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 2).
			Margin(1, 0)

	responseStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#04B575")).
			Padding(1, 2).
			Margin(1, 0)
)
