package models

import (
	"flean/tui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"slices"
	"strings"
)

type MultiOptionsModel struct {
	Question string
	Options  []string
	Headers  []string
	Selected []string
	OnSelect func(index int, option string) []string
	cursor   int
}

func (model MultiOptionsModel) Init() tea.Cmd {
	return nil
}

func (model MultiOptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if model.cursor > 0 {
				model.cursor--
			}
		case "down":
			if model.cursor < len(model.Options)-1 {
				model.cursor++
			}
		case "enter", " ":
			model.Selected = model.OnSelect(model.cursor, model.Options[model.cursor])
			return model, nil
		case "f":
			return model, tea.Quit
		case "q":
			os.Exit(0)
			return model, tea.Quit
		}
	}
	return model, nil
}

func (model MultiOptionsModel) View() string {
	style := lipgloss.NewStyle()
	msg := tui.LogStyle.Render(model.Question) + "\n"
	for i, v := range model.Options {
		cursor := "  "
		if slices.Contains(model.Selected, v) {
			if model.cursor == i {
				cursor = ">>"
			}
			msg += strings.TrimRight(tui.LogStyle.Bold(true).Render(fmt.Sprintf("%s %d. %s\n", cursor, i+1, v)), " ")
		} else if model.cursor == i {
			cursor = ">>"
			msg += strings.TrimRight(style.Foreground(tui.Purple).Render(fmt.Sprintf("%s %d. %s\n", cursor, i+1, v)), " ")
		} else {
			msg += fmt.Sprintf("%s %d. %s\n", cursor, i+1, v)
		}
	}
	msg += fmt.Sprintf("\nPress %s to finish\nPress %s to cancel project creation\n", style.Bold(true).Foreground(lipgloss.Color("#0FFF50")).Render("f"), style.Bold(true).Foreground(tui.Red).Render("q"))
	return style.BorderForeground(tui.Purple).BorderStyle(lipgloss.DoubleBorder()).Padding(1).Render(msg) + "\n"
}
