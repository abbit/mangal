package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/luevano/mangal/tui/base"
	"github.com/luevano/mangal/tui/model"
)

func Run(state base.State) error {
	program := tea.NewProgram(model.New(state), tea.WithAltScreen(), tea.WithMouseCellMotion())

	_, err := program.Run()
	return err
}
