package models

import (
	"flean/tui"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type LoadingModel struct {
	Message string
	Process func() (msg tea.Msg)
	spinner spinner.Model
}

func NewLoader(message string, process func() (msg tea.Msg)) LoadingModel {
	return LoadingModel{
		Message: message,
		Process: process,
		spinner: spinner.New(spinner.WithSpinner(spinner.Dot)),
	}
}

func (model LoadingModel) Init() tea.Cmd {
	return tea.Batch(model.spinner.Tick, func() tea.Msg {
		return model.Process()
	})
}

func (model LoadingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		model.spinner, cmd = model.spinner.Update(msg)
		return model, cmd
	case tea.QuitMsg:
		return model, tea.Quit
	default:
		if _, ok := msg.(tea.QuitMsg); ok {
			return model, tea.Quit
		}
		return model, nil
	}
}

func (model LoadingModel) View() string {
	return tui.LogStyle.Render(model.Message + " " + model.spinner.View())
}
