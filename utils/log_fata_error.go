package utils

import (
	"flean/tui"
	"github.com/charmbracelet/lipgloss"
	"log"
)

func LogFatalError(err error) {
	log.Fatal(lipgloss.NewStyle().Foreground(tui.Red).Render(err.Error()))
}
