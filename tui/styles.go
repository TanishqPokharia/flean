package tui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

var LogStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#40e0d0"))
var TableHeaderStyle lipgloss.Style = table.DefaultStyles().Header.BorderStyle(lipgloss.DoubleBorder())
var TableSelectedStyle lipgloss.Style = table.DefaultStyles().Selected.Foreground(lipgloss.Color("#40e0d0")).Blink(true)
var Purple = lipgloss.Color("#CF9FFF")
var Red = lipgloss.Color("#DE3163")
