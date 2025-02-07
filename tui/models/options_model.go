package models

import (
	"flean/tui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type OptionsModel struct {
	Question        string
	ProgressMessage func(index int) string
	Options         []string
	Index           int
	OnSelect        func(index int)
}

func NewOptionsModel(model OptionsModel) OptionsModel {
	return model
}

func (model OptionsModel) Init() tea.Cmd {
	return nil
}

func (model OptionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			fmt.Println("Aborting...")
			os.Exit(0)
			return model, tea.Quit
		case "up":
			if model.Index > 0 {
				model.Index--
			}
		case "down":
			if model.Index < len(model.Options)-1 {
				model.Index++
			}
		case "enter", " ":
			model.OnSelect(model.Index)
			return model, tea.Quit
		}
	case error:
		return model, nil
	}
	return model, nil

}

func (model OptionsModel) View() string {
	style := lipgloss.NewStyle()
	msg := tui.LogStyle.Render(model.Question + "\n")
	selectedOptionStyle := style.Foreground(tui.Purple).Bold(true)
	for i, option := range model.Options {
		cursor := "  "
		if model.Index == i {
			cursor = selectedOptionStyle.Blink(true).Render(">>")
		}
		if i == 0 {
			msg += "\n"
		}
		msg += fmt.Sprintf("%s %s\n", cursor, func() string {
			if model.Index == i {
				return selectedOptionStyle.Render(option)
			}
			return option
		}())
	}
	msg += fmt.Sprintf("\nPress %s to quit\n", style.Bold(true).Foreground(tui.Red).Render("q"))
	return msg
}
